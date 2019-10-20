#!/bin/sh

echo "srcGo/services/*"
go test ./srcGo/services -test.v -cover

echo ""
echo "srcGo/models/*"
GOOS=js GOARCH=wasm go test -exec="node $(go env GOROOT)/misc/wasm/wasm_exec" ./srcGo/models -test.v -cover
