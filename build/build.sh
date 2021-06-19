#!/bin/bash

VERSION=$(git describe --tags --always)
LDFLAGS="-s -w -X github.com/fenixsoft/fenix-cli/src/environments.Version=$VERSION"

go build -ldflags "${LDFLAGS}"
