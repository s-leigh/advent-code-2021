#!/bin/zsh

set -e

go test ./...
./gradlew build
