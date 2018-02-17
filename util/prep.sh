#! /bin/sh

set -xue

# Get latest foreman-cc package
cc_project="$(git ls-remote --get-url origin | sed 's/.*ghe.corp.cybergarage.org.//;s/-go\(.git\)\?$/-cc/')"
github_url="https://ghe.corp.cybergarage.org/api/v3/repos/$cc_project/releases"
deb_url=$(curl -sL "$github_url" | jq -r '.[0].assets[] | select(.name | contains(".deb")) | .browser_download_url')
curl -sL "$deb_url" > foreman-cc.deb
dpkg -i foreman-cc.deb
