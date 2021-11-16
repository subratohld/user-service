#!/bin/sh

PROJ_DIR="$(cd $(dirname $0)/.. && pwd)"

commitId=$(git rev-parse --short HEAD)

docker build . -t "subratohld/user-service:${commitId}"