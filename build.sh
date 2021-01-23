#!/bin/sh
cd ./src
go build -o ../bin/compose-generator compose-generator.go
cd ..