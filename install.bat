@echo off
mkdir bin 2> NUL
cd src
go install compose-generator.go
cd ..