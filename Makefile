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
GOPATH=$(shell pwd)

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
	${PACKAGE_ID}/rpc/graphite \
	${PACKAGE_ID}/rpc/json \
	${PACKAGE_ID}/test \
	${PACKAGE_ID}

SOURCE_DIR=src/${GITHUB}/foreman
BINARY_DEAMON=${GITHUB}/${DEAMON_NAME}
BINARY_TESTING=${GITHUB}/${TESTING_NAME}
BINARYIES=${BINARY_DEAMON} ${BINARY_TESTING}

CGO_LDFLAGS += -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog -luuid -lalglib
export CGO_LDFLAGS

.PHONY: version

all: test

VERSION_GO="./foreman/version.go"

${VERSION_GO}: ./foreman/version.gen
	$< > $@

version: ${VERSION_GO}

format:
	gofmt -w src/${GITHUB} ${PACKAGE_NAME} ${DEAMON_NAME} ${TESTING_NAME}

const: $(shell find ${SOURCE_DIR} -type f -name '*.csv')
	pushd ${SOURCE_DIR} && ./constants.go.gen > constants.go  && popd
	pushd ${SOURCE_DIR}/fql && ./constants.go.gen > constants.go  && popd
	pushd ${SOURCE_DIR}/action && ./constants.go.gen > constants.go  && popd
	pushd ${SOURCE_DIR}/rpc/json && ./constants.go.gen > constants.go  && popd
	pushd ${SOURCE_DIR}/errors && ./errors.go.gen > errors.go  && popd
	
antlr:
	pushd ${SOURCE_DIR}/fql && antlr4 -package fql -Dlanguage=Go FQL.g4 && popd

build: version const antlr format $(shell find ${SOURCE_DIR} -type f -name '*.go')
	go build -v ${PACKAGES}

test: build 
	go test -v -cover ${PACKAGES}

install: build
	go install ${BINARYIES}

clean:
	rm ${PREFIX}/bin/*
	rm -rf _obj
	go clean -i ${PACKAGES}
