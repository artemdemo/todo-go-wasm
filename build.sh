#!/bin/bash

GOOS=js GOARCH=wasm go build -o build/main.wasm srcGo/*.go
