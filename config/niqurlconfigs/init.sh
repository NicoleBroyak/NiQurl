#!/bin/sh
cd niqurl/cmd/server
go get -u
go mod tidy
go build -o server