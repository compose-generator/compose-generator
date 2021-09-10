#!/bin/sh
mkdir -p bin
cd src
env GOOS=linux GOARCH=amd64 go build -x -o ../bin/compose-generator compose-generator.go
cd ..
chmod -R 777 predefined-services