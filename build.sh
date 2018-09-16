#!/bin/bash

set -xe

REPO_PATH=github.com/fucangyu/mblog
GIT_SHA=$(git rev-parse --short HEAD || echo "GitNotFound")

# Set GO_LDFLAGS="-s" for building without symbols for debugging.
GO_LDFLAGS="-X ${REPO_PATH}/cmd.GitSHA=${GIT_SHA}"

go build -ldflags "$GO_LDFLAGS" -o mblog $REPO_PATH
