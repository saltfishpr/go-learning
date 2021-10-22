#!/bin/bash
echo "----- Build -----"
docker build --pull -f "$WORKDIR"/build/Dockerfile --tag "$APP_NAME":latest .
echo "-----  Run  -----"
docker run --rm "$APP_NAME"
