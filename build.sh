#/bin/bash

GOOS=js GOARCH=wasm go build -o ./main.wasm ../cmd/VersionA-wasm/