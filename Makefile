###################################################################
#
# foreman-go
#
# Copyright (C) Satoshi Konno 2017
#
# This is licensed under BSD-style license, see file COPYING.
#
###################################################################

# Need bash for pushd, popd
SHELL := bash

PREFIX?=$(shell pwd)
GOPATH:=$(shell pwd)
export GOPATH

GITHUB_ROOT=github.com/cybergarage

PACKAGE_NAME=foreman
DAEMON_NAME=foremand
TESTING_NAME=foremantest

GITHUB=${GITHUB_ROOT}/foreman-go

GO_GRAPHITE_GITHUB=${GITHUB_ROOT}/go-graphite
GO_GRAPHITE_PACKAGE_NAME=net/graphite
GO_GRAPHITE_GITHUB_ID=${GO_GRAPHITE_GITHUB}.git/${GO_GRAPHITE_PACKAGE_NAME}
GO_GRAPHITE_PACKAGE_ID=${GO_GRAPHITE_GITHUB}/${GO_GRAPHITE_PACKAGE_NAME}
GO_GRAPHITE_PACKAGES=${GO_GRAPHITE_PACKAGE_ID}

GITHUB_ID=${GITHUB}.git/${PACKAGE_NAME}
PACKAGE_ID=${GITHUB}/${PACKAGE_NAME}
PACKAGES=\
	${PACKAGE_ID} \
	${PACKAGE_ID}/logging \
	${PACKAGE_ID}/errors \
	${PACKAGE_ID}/registry \
	${PACKAGE_ID}/metric \
	${PACKAGE_ID}/register \
	${PACKAGE_ID}/action \
	${PACKAGE_ID}/kb \
	${PACKAGE_ID}/qos \
	${PACKAGE_ID}/fql \
	${PACKAGE_ID}/rpc \
	${PACKAGE_ID}/rpc/graphite \
	${PACKAGE_ID}/rpc/json \
	${PACKAGE_ID}/discovery \
	${PACKAGE_ID}/fd \
	${PACKAGE_ID}/test

SOURCE_DIR=src/${GITHUB}/foreman
BINARY_DAEMON=${GITHUB}/${DAEMON_NAME}
BINARY_TESTING=${GITHUB}/${TESTING_NAME}
BINARIES=${BINARY_DAEMON} ${BINARY_TESTING}

CGO_LDFLAGS += -lforeman++ -lalglib -lsqlite3 -luuid -lm
#CGO_LDFLAGS += -lfolly -lgflags -lglog

HAVE_PYTHON_CONFIG := $(shell command -v python-config 2> /dev/null)
all:
ifdef HAVE_PYTHON_CONFIG
    CGO_CFLAGS += $(shell python-config --includes)
    CGO_LDFLAGS += $(shell python-config --libs)
endif

HAVE_PKG_CONFIG := $(shell command -v pkg-config 2> /dev/null)
all:
ifdef HAVE_PYTHON_CONFIG
    CGO_CFLAGS += $(shell pkg-config lua --silence-errors --cflags)
    CGO_LDFLAGS += $(shell pkg-config lua --silence-errors --libs)
endif

export CGO_LDFLAGS

CONST_CSVS = $(wildcard $(SOURCE_DIR)/common/*.csv)
CONST_GENS = $(shell find $(SOURCE_DIR) -type f -name '*.go.gen')
CONST_GOS = $(basename $(CONST_GENS))

GO_FILES = $(shell find $(SOURCE_DIR) -type f -name '*.go')
ANTLR_FILES = $(addsuffix .go, $(addprefix $(SOURCE_DIR)/fql/fql_, base_listener lexer listener parser))

HAVE_ANTLR4 := $(shell which antlr4)
all:
ifdef HAVE_ANTLR4
    ANTLR=antlr4
else
    ANTLR=antlr
endif

.PHONY: version antlr clean

all: test

VERSION_GO=${SOURCE_DIR}/version.go

${VERSION_GO}: ${SOURCE_DIR}/version.gen
	$< > $@

version: ${VERSION_GO}

$(CONST_GOS):  $(CONST_GENS) $(CONST_CSVS)
	cd $(dir $@) && ./$(notdir $@).gen > $(notdir $@)

$(ANTLR_FILES): $(SOURCE_DIR)/fql/FQL.g4
	- cd ${SOURCE_DIR}/fql && ${ANTLR} -package fql -Dlanguage=Go FQL.g4

antlr: $(ANTLR_FILES)

const: $(CONST_GOS) antlr

format:
	gofmt -w src/${GITHUB} ${PACKAGE_NAME} ${DAEMON_NAME} ${TESTING_NAME}

const: $(shell find ${SOURCE_DIR} -type f -name '*.csv')
	pushd ${SOURCE_DIR} && ./constants.go.gen > constants.go  && popd
	pushd ${SOURCE_DIR}/fql && ./constants.go.gen > constants.go  && popd
	pushd ${SOURCE_DIR}/action && ./constants.go.gen > constants.go  && popd
	pushd ${SOURCE_DIR}/rpc && ./constants.go.gen > constants.go  && popd
	pushd ${SOURCE_DIR}/rpc/json && ./constants.go.gen > constants.go  && popd
	pushd ${SOURCE_DIR}/errors && ./errors.go.gen > errors.go  && popd
	
antlr:
	- pushd ${SOURCE_DIR}/fql && ${ANTLR} -package fql -Dlanguage=Go FQL.g4 && popd

vet: format
	go vet ${PACKAGES}

build: antlr vet
	go build -v ${PACKAGES}

test: antlr vet
	go test -v -cover -timeout 300s ${PACKAGES}

install: antlr vet
	go install -ldflags '${go_linker_flags}' -v ${BINARIES}

clean:
	-rm ${PREFIX}/bin/*
	rm -rf _obj
	go clean -i ${PACKAGES}
