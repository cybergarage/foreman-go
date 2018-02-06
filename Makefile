###################################################################
#
# foreman-go
#
# Copyright (C) Satoshi Konno 2017
#
# This is licensed under BSD-style license, see file COPYING.
#
###################################################################

PREFIX?=$(shell pwd)
GOPATH=$(shell pwd)

GITHUB_ROOT=github.com/cybergarage

PACKAGE_NAME=foreman
BINARY_NAME=foremand

GITHUB=${GITHUB_ROOT}/foreman-go

GO_GRAPHITE_GITHUB=${GITHUB_ROOT}/go-graphite
GO_GRAPHITE_PACKAGE_NAME=net/graphite
GO_GRAPHITE_GITHUB_ID=${GO_GRAPHITE_GITHUB}.git/${GO_GRAPHITE_PACKAGE_NAME}
GO_GRAPHITE_PACKAGE_ID=${GO_GRAPHITE_GITHUB}/${GO_GRAPHITE_PACKAGE_NAME}
GO_GRAPHITE_PACKAGES=${GO_GRAPHITE_PACKAGE_ID}

GITHUB_ID=${GITHUB}.git/${PACKAGE_NAME}
PACKAGE_ID=${GITHUB}/${PACKAGE_NAME}
PACKAGES=\
	${PACKAGE_ID}/log \
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
	${PACKAGE_ID}

SOURCE_DIR=src/${GITHUB}/foreman
BINARY_ID=${GITHUB}/${BINARY_NAME}
BINARYIES=${BINARY_ID}

.PHONY: version

all: test

VERSION_GO="./foreman/version.go"

${VERSION_GO}: ./foreman/version.gen
	$< > $@
	git commit -m "Updated version." $@

version: ${VERSION_GO}

format:
	gofmt -w src/${GITHUB} ${PACKAGE_NAME} ${BINARY_NAME}

const: $(shell find ${SOURCE_DIR} -type f -name '*.csv')
	pushd ${SOURCE_DIR} && ./constants.go.gen > constants.go  && popd
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
