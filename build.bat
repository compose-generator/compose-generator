@echo off
mkdir bin
cd src

go env -w GOOS=windows
go env -w GOARCH=arm
go env -w GOARM=7
go build -o ../bin/compose-generator-armv7.exe compose-generator.go

go env -w GOOS=windows
go env -w GOARCH=amd64
go build -o ../bin/compose-generator-amd64.exe compose-generator.go

cd ..