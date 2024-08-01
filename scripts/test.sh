#!/usr/bin/env sh

go generate
go test -cover -v ./...
