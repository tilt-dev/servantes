#!/usr/bin/env bash

docker build ./pingpong/ -t gcr.io/windmill-public-containers/servantes/pingpong && \
docker push gcr.io/windmill-public-containers/servantes/pingpong
