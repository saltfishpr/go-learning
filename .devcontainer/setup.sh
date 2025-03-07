#!/bin/sh

go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

git config --global http.https://github.com.proxy http://host.docker.internal:7897
git config --global https.https://github.com.proxy http://host.docker.internal:7897
