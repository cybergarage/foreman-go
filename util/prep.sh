#! /bin/sh

set -xue
mode=${1:-test}

yum install foreman-cc java-1.8.0-openjdk --enablerepo=apj-platform_rpms-"$mode" -y

if ! [ -d src/github.com/cybergarage/foreman-go ]
then
	mkdir -p src/github.com/cybergarage
	ln -sf "$PWD" src/github.com/cybergarage/foreman-go 
fi
