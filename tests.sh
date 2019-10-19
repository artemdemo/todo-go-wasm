#!/bin/sh

echo "srcGo/services/*"
go test ./services -test.v -cover
