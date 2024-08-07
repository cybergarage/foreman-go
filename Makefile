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

GITHUB_ROOT=github.com/cybergarage

PACKAGE_NAME=foreman
MODULE_NAME=foreman-go

DAEMON_NAME=foremand
TESTING_NAME=foremantest
TESTING_PROFILE_FILENAME=${TESTING_NAME}.prof

TESTING_PROFILE_FILENAME=foreman.test.prof


GO_GRAPHITE_GITHUB=${GITHUB_ROOT}/go-graphite
GO_GRAPHITE_PACKAGE_NAME=net/graphite
GO_GRAPHITE_GITHUB_ID=${GO_GRAPHITE_GITHUB}.git/${GO_GRAPHITE_PACKAGE_NAME}
GO_GRAPHITE_PACKAGE_ID=${GO_GRAPHITE_GITHUB}/${GO_GRAPHITE_PACKAGE_NAME}
GO_GRAPHITE_PACKAGES=${GO_GRAPHITE_PACKAGE_ID}

PACKAGE_ID=${GITHUB_ROOT}/${MODULE_NAME}/${PACKAGE_NAME}
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

SOURCE_DIR=${PACKAGE_NAME}

BINARY_DAEMON=${GITHUB_ROOT}/${MODULE_NAME}/${DAEMON_NAME}
BINARY_TESTING=${GITHUB_ROOT}/${MODULE_NAME}/${TESTING_NAME}
BINARIES=${BINARY_DAEMON} ${BINARY_TESTING}

CGO_CFLAGS =
CGO_LDFLAGS = -lforeman++ -lstdc++ -lalglib -lsqlite3 -luuid -lcurl -lm

# NOTE : Enable for folly
# CGO_LDFLAGS += -lfolly -lgflags -lglog

GCFLAGS="-N -l"

HAVE_PYTHON_CONFIG := $(shell command -v python-config 2> /dev/null)
all:
ifdef HAVE_PYTHON_CONFIG
    CGO_CFLAGS += $(shell python-config --includes)
    CGO_LDFLAGS += $(shell python-config --ldflags)
else
    HAVE_PYTHON3_CONFIG := $(shell command -v python3-config 2> /dev/null)
    all:
    ifdef HAVE_PYTHON3_CONFIG
        CGO_CFLAGS += $(shell python3-config --includes)
        CGO_LDFLAGS += $(shell python3-config --ldflags --embed)
    endif
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

HAVE_ANTLR4 := $(shell which antlr4)
all:
ifdef HAVE_ANTLR4
    ANTLR=antlr4
else
    ANTLR=antlr
endif

.PHONY: version clean

all: test

VERSION_GO=${SOURCE_DIR}/version.go

${VERSION_GO}: ${SOURCE_DIR}/version.gen
	$< > $@

version: ${VERSION_GO}

$(CONST_GOS):  $(CONST_GENS) $(CONST_CSVS)
	cd $(dir $@) && ./$(notdir $@).gen > $(notdir $@)

FQL_ANTLR_FILES = $(addsuffix .go, $(addprefix $(SOURCE_DIR)/fql/fql_, base_listener lexer listener parser)) $(SOURCE_DIR)/fql/FQL.g4

QOS_ANTLR_FILES = $(addsuffix .go, $(addprefix $(SOURCE_DIR)/kb/knowledge_, base_listener lexer listener parser)) $(SOURCE_DIR)/kb/Knowledge.g4

$(ANTLR_FILES): $(FQL_ANTLR_FILES) $(QOS_ANTLR_FILES)

antlr: $(ANTLR_FILES)
	- pushd ${SOURCE_DIR}/fql && ${ANTLR} -package fql -Dlanguage=Go FQL.g4 && popd
	- pushd ${SOURCE_DIR}/kb && ${ANTLR} -package kb -Dlanguage=Go Knowledge.g4 && popd

const: $(CONST_GOS) antlr

format:
	gofmt -w ${PACKAGE_NAME} ${DAEMON_NAME} ${TESTING_NAME}

const: $(shell find ${SOURCE_DIR} -type f -name '*.csv')
	pushd ${SOURCE_DIR} && ./constants.go.gen > constants.go && popd
	pushd ${SOURCE_DIR}/fql && ./constants.go.gen > constants.go && popd
	pushd ${SOURCE_DIR}/action && ./constants.go.gen > constants.go && popd
	pushd ${SOURCE_DIR}/rpc && ./constants.go.gen > constants.go && popd
	pushd ${SOURCE_DIR}/rpc/json && ./constants.go.gen > constants.go && popd
	pushd ${SOURCE_DIR}/errors && ./errors.go.gen > errors.go  && popd
	pushd ${SOURCE_DIR}/discovery && ./constants.go.gen > constants.go && popd
	pushd ${SOURCE_DIR}/logging && ./constants.go.gen > constants.go && popd
	
vet: format
	go vet ${PACKAGES}

build: antlr vet
	go build -v -gcflags=${GCFLAGS} ${PACKAGES}

test: antlr vet
	go test -v -cover -parallel 1 -timeout 1200s ${PACKAGES}

profile: 
	go test -v -cover -timeout 300s -memprofilerate 1 -memprofile ${TESTING_PROFILE_FILENAME} ${PACKAGE_ID}
	go tool pprof ${TESTING_PROFILE_FILENAME}
	rm ${TESTING_PROFILE_FILENAME}

install: antlr vet
	go install -v -gcflags=${GCFLAGS} ${BINARIES}

clean:
	-rm ${PREFIX}/bin/*
	rm -rf _obj
	go clean -i ${PACKAGES}
