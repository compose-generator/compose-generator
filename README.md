<p align="center">
  <img alt="Compose Generator Logo" src="./docs/docs/static/avatar.png" height="220" />
  <h3 align="center">Compose Generator</h3>
  <p align="center">Easy to use cli to generate Docker Compose YAML configuration files.</p>
  <p align="center">
    <a target="_blank" href="https://github.com/compose-generator/compose-generator/releases/latest"><img src="https://img.shields.io/github/v/release/compose-generator/compose-generator?include_prereleases"></a>
    <a target="_blank" href="https://hub.docker.com/r/chillibits/compose-generator"><img src="https://img.shields.io/docker/pulls/chillibits/compose-generator"></a>
    <a target="_blank" href="./.github/workflows/ci.yml"><img src="https://github.com/compose-generator/compose-generator/workflows/Go%20CI/badge.svg"></a>
    <a target="_blank" href="./.github/workflows/codeql-analysis.yml"><img src="https://github.com/compose-generator/compose-generator/actions/workflows/codeql-analysis.yml/badge.svg"></a>
    <a target="_blank" href="https://goreportcard.com/report/github.com/compose-generator/compose-generator"><img src="https://goreportcard.com/badge/github.com/compose-generator/compose-generator"></a>
    <a href="https://codecov.io/gh/compose-generator/compose-generator"><img src="https://codecov.io/gh/compose-generator/compose-generator/branch/main/graph/badge.svg?token=r9pWf0GCXg"/></a>
    <a target="_blank" href="https://makeapullrequest.com"><img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg"></a>
    <a target="_blank" href="./LICENSE.md"><img src="https://img.shields.io/github/license/compose-generator/compose-generator"></a>
  </p>
</p>

---

## Documentation
Please visit the documentation on [compose-generator.com](https://www.compose-generator.com).

## Usage
You can use the Compose Generator CLI by directly installing it on your Docker host system, install it via npm or by generating your compose file with the Compose Generator Docker container on the fly.

### Install Compose Generator CLI
To install Compose Generator on your system, please visit the [installation section](https://www.compose-generator.com/install/linux/) in the documentation. Compose Generator is available for the latest versions of Alpine, CentOS, Debian, Fedora, Raspbian, Ubuntu, Windows. If you want to install Compose Generator manually, please look at the table below.

## QuickStart with Docker
*Note for Windows users: This command does not work with Windows CMD command line. Please use Windows PowerShell instead.*

```sh
$ docker run --rm -it -v ${pwd}:/cg/out chillibits/compose-generator [<command>]
```

## Supported host systems & file downloads
There are also downloadable packages available for all supported platforms:

| **Platform**                | **x86_64 / amd64**                                                                     | **i386**                                                                             | **armv5**                                                                              | **armv6**                                                                              | **armv7**                                                                              | **arm64**                                                                              |
|-----------------------------|----------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------|
| **Darwin / MacOS (tar.gz)** | [download](../../releases/download/0.8.0/compose-generator_0.8.0_darwin_amd64.tar.gz)  | -                                                                                    | -                                                                                      | -                                                                                      | -                                                                                      | -                                                                                      |
| **FreeBSD (tag.gz)**        | [download](../../releases/download/0.8.0/compose-generator_0.8.0_freebsd_amd64.tar.gz) | [download](../../releases/download/0.8.0/compose-generator_0.8.0_freebsd_386.tar.gz) | [download](../../releases/download/0.8.0/compose-generator_0.8.0_freebsd_armv5.tar.gz) | [download](../../releases/download/0.8.0/compose-generator_0.8.0_freebsd_armv6.tar.gz) | [download](../../releases/download/0.8.0/compose-generator_0.8.0_freebsd_armv7.tar.gz) | [download](../../releases/download/0.8.0/compose-generator_0.8.0_freebsd_arm64.tar.gz) |
| **Alpine (apk)**            | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_amd64.apk)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_386.apk)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_armv5.apk)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_armv6.apk)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_armv7.apk)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_arm64.apk)      |
| **CentOS (rpm)**            | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_amd64.rpm)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_386.rpm)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_armv5.rpm)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_armv6.rpm)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_armv7.rpm)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_arm64.rpm)      |
| **Debian (deb)**            | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_amd64.deb)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_386.deb)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_armv5.deb)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_armv6.deb)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_armv7.deb)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_arm64.deb)      |
| **Fedora (rpm)**            | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_amd64.rpm)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_386.rpm)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_armv5.rpm)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_armv6.rpm)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_armv7.rpm)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_arm64.rpm)      |
| **Raspbian (deb)**          | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_amd64.deb)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_386.deb)      | -                                                                                      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_armv6.deb)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_armv7.deb)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_arm64.deb)      |
| **Ubuntu (deb)**            | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_amd64.deb)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_386.deb)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_armv5.deb)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_armv6.deb)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_armv7.deb)      | [download](../../releases/download/0.8.0/compose-generator_0.8.0_linux_arm64.deb)      |
| **Windows Installer (msi)** | [download](../../releases/download/0.8.0/compose-generator_0.8.0_x64_setup.msi)        | [download](../../releases/download/0.8.0/compose-generator_0.8.0_x86_setup.msi)      | -                                                                                      | -                                                                                      | -                                                                                      | -                                                                                      |
| **Windows Portable (zip)**  | [download](../../releases/download/0.8.0/compose-generator_0.8.0_windows_amd64.zip)    | [download](../../releases/download/0.8.0/compose-generator_0.8.0_windows_386.zip)    | -                                                                                      | -                                                                                      | -                                                                                      | -                                                                                      |

## Contribute by providing predefined templates
If you miss a predefined template and you want to create one for the public, please read the [instructions to create one](./predefined-services/README.md). Fork the repository, create the template and open a pr to the `dev` branch.
The community is thankful for every predefined template!

## Contribute otherwise to the project
If you want to contribute to this project, please ensure you comply with the [contribution guidelines](CONTRIBUTING.md).

© Marc Auberer 2021