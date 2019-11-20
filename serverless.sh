#!/usr/bin/env bash

SRCPATH=$PWD/go/src/github.com/takeon-validation-wrangler
GOPATH=$PWD/go
cd $SRCPATH
go get -t -v ./...
env GOOS=linux go build  -o bin/main
serverless package --package pkg
serverless deploy --verbose
