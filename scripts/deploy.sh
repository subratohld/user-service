#!/bin/sh

PROJ_DIR="$(cd $(dirname $0)/.. && pwd)"

commitId=$(git rev-parse --short HEAD)

cp "${PROJ_DIR}/deployments/deployment.yaml" "${PROJ_DIR}/deployments/temp-deployment.yaml"

sed -i -e "s/VERSION/${commitId}/g" "${PROJ_DIR}/deployments/temp-deployment.yaml"

kubectl apply -f "${PROJ_DIR}/deployments/temp-deployment.yaml"