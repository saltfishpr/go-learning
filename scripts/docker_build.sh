#!/bin/bash
echo "----- Build -----"
docker build --pull -f "$WORKDIR"/build/Dockerfile --tag "$APP_NAME":latest .
echo "-----  Run  -----"
docker run --rm -p 8090:8090 -p 8091:8091 "$APP_NAME"
