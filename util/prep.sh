#! /bin/sh
# Assume foreman-cc tags are in sync
tag=$(git describe --abbrev=0 --tags)
curl -sLO "https://https://raw.github.com/cybergarage/foreman-cc/releases/download/$tag/foreman-cc_$tag-1_amd64.deb"
dpkg -i "foreman-cc_$tag-1_amd64.deb"
