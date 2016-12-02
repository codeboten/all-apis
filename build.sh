#!/bin/bash

echo "Building datapi"
docker build -f datapi/Dockerfile -t alex/datapi datapi

echo "Building netapi"
cd netapi/app
GOOS=linux GOARCH=amd64 go build -o server
cd -
docker build -f netapi/Dockerfile -t alex/netapi netapi
