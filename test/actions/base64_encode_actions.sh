#!/bin/bash
ls -1 *.py | xargs -Ifname sh -c 'base64 fname > fname.base64'
