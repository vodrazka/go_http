#!/bin/bash

curl --header "Content-Type: application/json" \
  --request POST \
  --data 'test from curl' \
  http://localhost:8080/
