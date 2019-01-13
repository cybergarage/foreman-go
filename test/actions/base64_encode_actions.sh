#!/bin/bash
ls -1 *.py | xargs -Ifname sh -c 'base64 -w 0 fname > fname.base64'
