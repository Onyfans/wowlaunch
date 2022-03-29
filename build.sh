#!/usr/bin/env sh

go build
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build

tar czf wowlaunch-linux-amd64.tar.gz wowlaunch config.yaml
zip wowlaunch-windows-amd64 wowlaunch.exe config.yaml
