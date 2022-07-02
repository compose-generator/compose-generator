---
title: Install on Windows
---

### Install via winget
Compose Generator gets distributed for Windows via the new Windows package manager CLI called [winget](https://github.com/microsoft/winget-cli).

As soon as the Windows package manager is installed on your Windows machine, you can open up powershell and execute the following installation command: <br>
```sh
winget install ChilliBits.CCom
winget install ChilliBits.ComposeGenerator
```
After installing Compose Generator, you should be able to run Compose Generator. If not, please restart your powershell instance to make it reload the available commands.

### Install via installer

[Download 64bit](https://github.com/compose-generator/compose-generator/releases/latest/download/compose-generator_x64_setup.msi){ .md-button .md-button--primary .md-button--small }
[Download 32bit](https://github.com/compose-generator/compose-generator/releases/latest/download/compose-generator_x86_setup.msi){ .md-button .md-button--primary .md-button--small }

Note: It is possible, that you see an error notification by your anti-virus software, that Compose Generator is a potentially dangerous application. This can happen if the latest version was released very recently and too few people downloaded it already. If you want to validate the file signature, you can visit [this GitHub repo](https://github.com/microsoft/winget-pkgs/tree/master/manifests/c/ChilliBits/ComposeGenerator), select your version and open the file `ChilliBits.ComposeGenerator.installer.yaml`. There you can find the SHA256 signatures for the `x64` and `x86` exe installers.

### Use
```sh
compose-generator [<command>]
```

*[CLI]: Command Line Interface
