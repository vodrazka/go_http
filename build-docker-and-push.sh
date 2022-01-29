#!/bin/bash

docker build . -t vodrazka/go_http
docker push vodrazka/go_http
