#!/bin/bash

docker run -p 3000:3000 --name datapi -d -v $PWD/datapi/app:/app alex/datapi:latest

docker run -p 5000:5000 --name netapi -d -v $PWD/netapi/app:/app alex/netapi:latest
