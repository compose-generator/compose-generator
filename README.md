# Compose Generator
![Build passing](https://github.com/marcauberer/compose-generator/workflows/Go%20CI/badge.svg)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

## Usage
You can use the Compose Generator CLI by directly installing it on your Docker host system or by generating your compose file with the Compose Generator Docker container.

### Install Compose Generator CLI
```sh
$ apt-get install compose-generator
```

## Generate compose file on the fly with Docker container
```sh
$ docker run --rm -it -v .:/compose-generator/out ghcr.io/marcauberer/compose-generator
```

## Supported host systems


## Contribute to the project
If you want to contribute to this project, please feel free to send us a pull request.

© Marc Auberer 2021