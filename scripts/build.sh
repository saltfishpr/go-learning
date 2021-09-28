#!/bin/bash
VERSION=$(git --git-dir=$PWD/.git describe --tags --always)

set -o errexit
set -o nounset
set -o pipefail

mkdir -p $PWD/output/bin/ && go build -ldflags "-X main.Version=$VERSION" -o $PWD/output/bin/$APP_NAME ./...
echo "#!/bin/bash" >$PWD/output/run.sh
echo "${PWD}/output/bin/${APP_NAME}" >>$PWD/output/run.sh
