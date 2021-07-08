#!/bin/sh
mkdir -p bin
cd src
env GOOS=linux GOARCH=arm GOARM=7 go build -x -o ../bin/compose-generator-armv7 compose-generator.go
env GOOS=linux GOARCH=amd64 go build -x -o ../bin/compose-generator-amd64 compose-generator.go
cd ..
chmod -R 777 predefined-services