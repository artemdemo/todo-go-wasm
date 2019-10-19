#!/bin/sh

echo "srcGo/services/*"
go test ./srcGo/services -test.v -cover
