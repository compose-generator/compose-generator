@echo off
cd src
go build -o ../bin/compose-generator.exe compose-generator.go
cd ..