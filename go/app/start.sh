#!/bin/bash -eu
go env
go mod tidy
go get github.com/codegangsta/gin
gin --immediate run main.go 0.0.0.0:8080
# go run main.go 0.0.0.0:8080
