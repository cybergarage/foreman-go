#! /bin/sh

set -xue

yum install foreman-cc --enablerepo=apj-platform_rpms-test -y

if ! [ -d src/github.com/cybergarage/foreman-go ]
then
	mkdir -p src/github.com/cybergarage
	ln -sf "$PWD" src/github.com/cybergarage/foreman-go 
fi
