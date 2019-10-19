#!/bin/sh

echo "srcGo/services/classnames"
go test srcGo/services/classnames_test.go srcGo/services/classnames.go -cover
