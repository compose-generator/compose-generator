<p align="center">
  <img alt="Compose Generator Logo" src="https://github.com/compose-generator/compose-generator/raw/main/media/logo-wide.png" height="280" />
  <h3 align="center">Compose Generator</h3>
  <p align="center">Easy to use cli tool to generate Docker Compose configurations blazingly fast.</p>
  <p align="center">
    <a target="_blank" href="https://github.com/compose-generator/compose-generator/releases/latest"><img src="https://img.shields.io/github/v/release/compose-generator/compose-generator?include_prereleases"></a>
    <a target="_blank" href="https://hub.docker.com/r/chillibits/compose-generator"><img src="https://img.shields.io/docker/pulls/chillibits/compose-generator"></a>
    <a target="_blank" href="./.github/workflows/ci.yml"><img src="https://github.com/compose-generator/compose-generator/workflows/Go%20CI/badge.svg"></a>
    <a target="_blank" href="./.github/workflows/codeql-analysis.yml"><img src="https://github.com/compose-generator/compose-generator/actions/workflows/codeql-analysis.yml/badge.svg"></a>
    <a target="_blank" href="https://goreportcard.com/report/github.com/compose-generator/compose-generator"><img src="https://goreportcard.com/badge/github.com/compose-generator/compose-generator"></a>
    <a target="_blank" href="https://makeapullrequest.com"><img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg"></a>
    <a target="_blank" href="./LICENSE.md"><img src="https://img.shields.io/github/license/compose-generator/compose-generator"></a>
  </p>
</p>

---

## Quick reference
- **Maintained by:** [the Compose Generator contributors](https://github.com/compose-generator/compose-generator)
- **Where to get help:** [Official Website](https://www.compose-generator.com), [GitHub](https://github.com/compose-generator/compose-generator)

## Supported tags and respective `Dockerfile` links
- `1.5.0`, `1.5`, `1`, `latest`
- `1.4.0`, `1.4`
- `1.3.0`, `1.3`
- `1.2.1`, `1.2`
- `1.1.0`, `1.1`
- `1.0.0`

## Quick reference (cont.)
- **Where to file issues:** https://github.com/compose-generator/compose-generator
- **Supported architectures:** `amd64`, `i386`, `arm32v6`, `arm32v7`
- **Image updates:** [releases page](https://github.com/compose-generator/compose-generator/releases)

## Documentation
Please visit the documentation on [compose-generator.com](https://www.compose-generator.com).

## Usage
You can use the Compose Generator CLI by directly installing it on your Docker host system<!--, install it via npm--> or by generating your compose file with the Compose Generator Docker container on the fly.

### Install Compose Generator CLI
For installation instructions for <!--NPM, -->Linux, Windows, etc., please visit the [installation guide](https://www.compose-generator.com/install/linux).

## QuickStart with Docker
**For Linux:**
```sh
$ docker run --rm -it -v /var/run/docker.sock:/var/run/docker.sock -v $(pwd):/cg/out chillibits/compose-generator [<command>]
```

**For Windows:**
```sh
$ docker run --rm -it -v /var/run/docker.sock:/var/run/docker.sock -v ${pwd}:/cg/out chillibits/compose-generator [<command>]
```
*Note: This command does not work with Windows CMD command line. Please use Windows PowerShell instead.*

## Contribute by providing predefined service templates
If you miss a predefined service and you want to create one for the public, please read the [instructions to create one](https://github.com/compose-generator/compose-generator/blob/main/predefined-services/README.md). Fork the repository, create the template and open a pr.
The community is thankful for every predefined template!

## Contribute otherwise to the project
If you want to contribute to this project, please ensure you comply with the [contribution guidelines](https://github.com/compose-generator/compose-generator/blob/main/CONTRIBUTING.md).

© Marc Auberer 2021