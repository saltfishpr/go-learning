#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

go mod tidy
mkdir -p "$WORKDIR"/output/bin/ && go build -ldflags "-X 'main.buildTag=$(date '+%F %H:%M:%S')'" -o "$WORKDIR"/output/bin/"$APP_NAME" "$APP_NAME"/cmd
echo "#!/bin/bash" >"$WORKDIR"/output/run.sh
echo "${WORKDIR}/output/bin/${APP_NAME}" >>"$WORKDIR"/output/run.sh
