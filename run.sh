#!/bin/bash

docker run -p 3000:3000 --name datapi -d -v $PWD/datapi/app:/app alex/datapi:latest
