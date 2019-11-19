#!/usr/bin/env bash

env GOOS=linux go build  -o bin/main
cd wrangler-deploy-repository
echo Packaging serverless bundle...
serverless package --package pkg
echo Deploying to AWS...
serverless deploy --verbose;
