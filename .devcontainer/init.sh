#!/bin/bash
set -o nounset
set -o errexit
set -o xtrace

echo "# 默认注释了源码镜像以提高 apt update 速度，如有需要可自行取消注释
deb https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye main contrib non-free
# deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye main contrib non-free
deb https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-updates main contrib non-free
# deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-updates main contrib non-free

deb https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-backports main contrib non-free
# deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-backports main contrib non-free

deb https://mirrors.tuna.tsinghua.edu.cn/debian-security bullseye-security main contrib non-free
# deb-src https://mirrors.tuna.tsinghua.edu.cn/debian-security bullseye-security main contrib non-free
" | sudo tee /etc/apt/sources.list > /dev/null

go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

go install github.com/go-task/task/v3/cmd/task@latest # task runner
go install github.com/segmentio/golines@latest # formatter
go install github.com/bufbuild/buf/cmd/buf@latest # buf

# download protoc to /usr/local/bin and google well known types to /usr/local/include
TMP="/tmp" && BIN="/usr/local/bin" && VERSION=3.20.3 && \
    curl -L https://github.com/protocolbuffers/protobuf/releases/download/v$VERSION/protoc-$VERSION-$(uname -s)-$(uname -m).zip -o $TMP/protoc.zip && \
    unzip -o $TMP/protoc.zip -d $TMP/protoc && \
    sudo mv $TMP/protoc/bin/protoc $BIN/protoc && \
    sudo mv $TMP/protoc/include/google /usr/local/include/google && \
    sudo chmod +x $BIN/protoc && \
    rm -rf $TMP/protoc.zip $TMP/protoc
