#!/bin/sh

PROJ_DIR="$(cd $(dirname $0)/.. && pwd)"

coverageFile="${PROJ_DIR}/share/cover.out"

go test -v -coverprofile="${coverageFile}" -covermode=count  ./internal/...
go tool cover -html="${coverageFile}" -o "${PROJ_DIR}/share/cover.html"