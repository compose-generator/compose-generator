@echo off
mkdir bin 2> NUL
cd src

go env -w GOOS=windows
go env -w GOARCH=amd64
go build -o ../bin/compose-generator.exe compose-generator.go

cd ..