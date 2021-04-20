---
title: Install on Windows
---

### Install via winget cli
Compose Generator gets distributed for Windows via the new Windows package manager called [winget](https://github.com/microsoft/winget-cli). <br>
In the future, winget will be available for download in the Microsoft Store. Currently, the easiest way to install winget is, to download it manually from GitHub. Visit the [installation instruction](https://github.com/microsoft/winget-cli#installing-the-client) from Microsoft.

As soon as the Windows package manager is installed on your Windows machine, you can open up powershell and execute the following installation command: <br>
```sh
winget install ChilliBits.ComposeGenerator
```
After installing Compose Generator, you should be able to run Compose Generator. If not, please restart your powershell instance to make it reload the available commands.

### Install via installer
To install Compose Generator on Windows, download the `ComposeGenerator_<version>_<arch>_Setup.exe` file [here](https://github.com/compose-generator/compose-generator/releases/latest) and run it. <br>
It is possible, that you see an error notification by your anti-virus software, that Compose Generator is a potentally dangerous application. This can happen if the latest version was released very recently and too few people downloaded it already. If you want to validate the file signature, you can visit [this GitHub repo](https://github.com/microsoft/winget-pkgs/tree/master/manifests/c/ChilliBits/ComposeGenerator), select your version and open the file `ChilliBits.ComposeGenerator.installer.yaml`. There you can find the SHA256 signatures for the `x64` and `x86` exe installers.

### Use
```sh
compose-generator
```