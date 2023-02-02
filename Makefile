SHELL := /bin/bash

.PHONY: build clean format pre-deploy deploy-all deploy-production deploy-staging test

build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/main main.go

clean:
	rm -rf ./bin

deploy-all: deploy-staging deploy-production

deploy-staging: pre-deploy
	serverless deploy  --verbose --stage staging --region us-east-2 --org serverlessorgname

deploy-production: pre-deploy
	serverless deploy  --verbose --stage production --region us-east-2 --org serverlessorgname

format:
	gofmt -s -w -l .

pre-deploy: clean build

test:
	source .env && go test ./...
