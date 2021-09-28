#!/bin/bash
docker build --build-arg APP_NAME=$APP_NAME --no-cache --pull -f $PWD/build/Dockerfile --tag $APP_NAME:latest .
docker run -d $APP_NAME
