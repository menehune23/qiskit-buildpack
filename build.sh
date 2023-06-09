#!/usr/bin/env bash

set -euo pipefail

mkdir -p out
GOOS=linux go build -o out/bin/detect ./cmd/main
cp out/bin/detect out/bin/build
cp buildpack.toml out/buildpack.toml
