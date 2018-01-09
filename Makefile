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
PACKAGES=${PACKAGE_ID} ${PACKAGE_ID}/log ${PACKAGE_ID}/errors ${PACKAGE_ID}/registry ${PACKAGE_ID}/metric ${PACKAGE_ID}/register ${PACKAGE_ID}/action ${PACKAGE_ID}/qos ${PACKAGE_ID}/fql

BINARY_ID=${GITHUB}/${BINARY_NAME}
BINARYIES=${BINARY_ID}

.PHONY: setup

all: test

VERSION_GO="./foreman/version.go"

${VERSION_GO}: ./foreman/version.gen
	$< > $@

version: ${VERSION_GO}

SETUP_CMD="./setup"

setup:
	@echo "#!/bin/bash" > ${SETUP_CMD}
	@echo "export GOPATH=\`pwd\`" >> ${SETUP_CMD}
	@echo "git pull" >> ${SETUP_CMD}
	@echo "mkdir -p src" >> ${SETUP_CMD}
	@echo "rm -rf src/${GITHUB_ROOT}" >> ${SETUP_CMD}
	@echo "# go-graphite" >> ${SETUP_CMD}
	@echo "go get -u ${GO_GRAPHITE_GITHUB_ID}" >> ${SETUP_CMD}
	@echo "pushd src && mv ${GO_GRAPHITE_GITHUB}.git ${GO_GRAPHITE_GITHUB} && popd" >> ${SETUP_CMD}
	@echo "# foreman-go" >> ${SETUP_CMD}
	@echo "go get -u ${GITHUB_ID}" >> ${SETUP_CMD}
	@echo "pushd src && mv ${GITHUB}.git ${GITHUB} && popd" >> ${SETUP_CMD}
	@chmod a+x ${SETUP_CMD}
	@./${SETUP_CMD}

format:
	gofmt -w src/${GITHUB} ${PACKAGE_NAME} ${BINARY_NAME}

build: format $(shell find . -type f -name '*.go')
	go build -v ${PACKAGES}

test: build 
	go test -v -cover ${PACKAGES}

install: build
	go install ${BINARYIES}

clean:
	rm ${PREFIX}/bin/*
	rm -rf _obj
	go clean -i ${PACKAGES}
