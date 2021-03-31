#!/bin/sh
sudo chown -R $(id -u):$(id -g) predefined-services
mkdir -p bin
cd src
go install compose-generator
cd ..
chmod -R 777 predefined-services
chmod -R 777 toolbox