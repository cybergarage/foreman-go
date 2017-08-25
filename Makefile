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

GITHUB=github.com/cybergarage
PACKAGES=${GITHUB}/foreman-go/foreman
DEP_PACKAGES=${GITHUB}/go-graphite/net/graphite

.PHONY: setup

VERSION_GO="./foreman/version.go"

${VERSION_GO}: ./foreman/version.gen
	$< > $@

version: ${VERSION_GO}

SETUP_CMD="./setup"

setup:
	@echo "#!/bin/bash" > ${SETUP_CMD}
	@echo "export GOPATH=\`pwd\`" >> ${SETUP_CMD}
	@echo "git pull" >> ${SETUP_CMD}
	@echo "go get -u ${PACKAGES}" >> ${SETUP_CMD}
	@echo "go get -u ${DEP_PACKAGES}" >> ${SETUP_CMD}
	@chmod a+x ${SETUP_CMD}
	@./${SETUP_CMD}

commit:
	pushd src/${GITHUB} && git commit -a && popd

push:
	pushd src/${GITHUB} && git push && popd

pull:
	pushd src/${GITHUB} && git pull && popd

diff:
	pushd src/${GITHUB} && git diff && popd

format:
	gofmt -w src/${GITHUB} foreman

package: format $(shell find . -type f -name '*.go')
	go build -v ${PACKAGES}

test: package
	go test -v -cover ${PACKAGES}

install: build
	go install ${PACKAGES}

clean:
	rm ${PREFIX}/bin/*
	rm -rf _obj
	go clean -i ${PACKAGES}
