# Compose Generator
![Build passing](https://github.com/marcauberer/compose-generator/workflows/Go%20CI/badge.svg)
[![Go Report](https://goreportcard.com/badge/github.com/marcauberer/compose-generator)](https://goreportcard.com/report/github.com/marcauberer/compose-generator)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

## Usage
You can use the Compose Generator CLI by directly installing it on your Docker host system or by generating your compose file with the Compose Generator Docker container.

### Install Compose Generator CLI
**Linux (Debian based distributions)**<br>
```sh
$ sudo apt-get update

$ 

$ sudo apt-get install compose-generator
```

**Linux (RPM based distributions)**<br>
```sh
$ sudo apt-get update

$

$ sudo apt-get install compose-generator
```

**Windows**
Compose Generator gets distributed for Windows via the new Windows package manager called [winget](https://github.com/microsoft/winget-cli). In the future, winget will be available for download in the Microsoft Store. Currently, the easiest way to install winget is, to download it manually from GitHub. Visit the [installation instruction](https://github.com/microsoft/winget-cli#installing-the-client) from Microsoft.

As soon as the Windows package manager is installed on your Windows machine, you can open powershell and execute this installation command:
```
$ winget install ChilliBits.ComposeGenerator
```

### Generate compose file on the fly with Docker container
*Note: This command does not work with Windows CMD command line. Please use PowerShell.*

```sh
$ docker run --rm -it -v ${pwd}:/compose-generator/out chillibits/compose-generator
```

## Supported host systems


## Contribute to the project
If you want to contribute to this project, please feel free to send us a pull request.

© Marc Auberer 2021