#!/bin/bash
# Linux
#ls -1 *.py | xargs -Ifname sh -c 'base64 -w 0 fname > fname.base64'
# MacOSX
ls -1 *.py | xargs -Ifname base64 -b 0 -i fname -o fname.base64
