#! /bin/sh

set -xue

# Get latest foreman-cc package
cc_project="$(git ls-remote --get-url origin | sed 's/.*ghe.corp.cybergarage.org.//;s/-go\(.git\)\?$/-cc/')"
github_url="https://ghe.corp.cybergarage.org/api/v3/repos/$cc_project/releases"
rpm_url=$(curl -sL "$github_url" | jq -r '.[0].assets[] | select(.name | contains(".rpm")) | .browser_download_url')

yum install "$rpm_url" -y

if ! [ -d src/github.com/cybergarage/foreman-go ]
then
	mkdir -p src/github.com/cybergarage
	ln -sf "$PWD" src/github.com/cybergarage/foreman-go 
fi
