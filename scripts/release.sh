#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

go mod tidy
mkdir -p "$PWD"/output/bin/ && go build -ldflags "-X learning/config.BuildDate=$(date '+%F_%H:%M:%S') -X learning/config.Release=true" -o "$PWD"/output/bin/"$APP_NAME" .
echo "#!/bin/bash" >"$PWD"/output/run.sh
echo "${PWD}/output/bin/${APP_NAME}" >>"$PWD"/output/run.sh
