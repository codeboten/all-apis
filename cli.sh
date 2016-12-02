#!/bin/bash

API_URL=localhost

function services() {
  curl ${API_URL}:8001/apis | jq .[].api
}

function doit() {
  api=$1
  action=$2

  method=GET
  endpoint=networks

  curl -X ${method} ${API_URL}:8000/${api}/${endpoint}
}

doit
