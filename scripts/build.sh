#!/bin/sh

PROJ_DIR="$(cd $(dirname $0)/.. && pwd)"

commitId=$(git rev-parse --short HEAD)

serviceName="user-service"

docker build . -t "subratohld/${serviceName}:${commitId}"