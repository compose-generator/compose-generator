#!/bin/sh
cd ./src
env GOOS=linux GOARCH=arm GOARM=7 go build -o ../bin/compose-generator-arm-v7 compose-generator.go
env GOOS=linux GOARCH=amd64 go build -o ../bin/compose-generator-amd64 compose-generator.go
cd ..