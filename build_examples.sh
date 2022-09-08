#!/bin/bash

mkdir -p bin

echo "Building to bin/ folder"

CGO_ENABLED=0 go build -v -o bin/ ./...
