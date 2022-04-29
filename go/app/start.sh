#!/bin/bash

# change line setting from CRLF to LF if neccessary 

go mod tidy
go install github.com/codegangsta/gin@latest
gin --immediate run main.go 0.0.0.0:8080
