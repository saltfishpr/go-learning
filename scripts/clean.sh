#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

if [ -d output ]; then
  echo "----- Clean Build  -----"
  rm -rf output
  echo "remove output"
else
  echo "no output to clean"
fi

if docker images | grep -q "$APP_NAME"; then
  echo "----- Clean Docker -----"
  docker rmi "$APP_NAME"
fi

if docker images | grep -q none; then
  echo "----- Prune Images -----"
  docker image prune -f
fi
