#!/bin/bash -eu
go env
go mod tidy
go install github.com/codegangsta/gin@latest
gin --immediate run main.go 0.0.0.0:8080
# go run main.go 0.0.0.0:8080
