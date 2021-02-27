---
title: Install on Windows
---

### Install via winget cli
Compose Generator gets distributed for Windows via the new Windows package manager called [winget](https://github.com/microsoft/winget-cli). <br>
In the future, winget will be available for download in the Microsoft Store. Currently, the easiest way to install winget is, to download it manually from GitHub. Visit the [installation instruction](https://github.com/microsoft/winget-cli#installing-the-client) from Microsoft.

As soon as the Windows package manager is installed on your Windows machine, you can open powershell and execute this installation command: <br>
```sh
winget install ChilliBits.ComposeGenerator
```
After installing Compose Generator, you should restart your powershell instance to make it reload the available commands.

### Install with installer
To install Compose Generator on Windows, download the `ComposeGenerator_<version>_<arch>_Setup.exe` file [here](https://github.com/compose-generator/compose-generator/releases/latest) and run it.

### Use
```sh
compose-generator
```