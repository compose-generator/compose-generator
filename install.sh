#!/bin/sh
sudo chown -R $(id -u):$(id -g) predefined-services
mkdir -p bin
cd src
go install compose-generator
cd ..