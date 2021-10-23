---
title: Build or install from source
---

### Setup
Before being able to compile the source code you first have to download it from GitHub and install Golang (if you don't have it already).

#### Clone source from GitHub
```sh
git clone https://github.com/compose-generator/compose-generator.git
cd compose-generator
```

#### Install Go
For compiling Compose Generator from source you need to have Go 1.16 or higher installed on your system. Please refer to the [official guide](https://golang.org/doc/install) on how to install Go.

### Build from source
If you want to build the executable, you can use the `build.sh` for Linux or the `build.bat` file for Windows:

=== "Linux"
    ```sh
    ./build.sh
    ```
=== "Windows"
    ```sh
    .\build.bat
    ```

The executables will end up in the `bin` subfolder for the `amd64` architecture.

### Install from source
If you don't care about the executable and you just want to use Compose Generator, you can use the `install.sh` for Linux or the `install.bat` for Windows. This will install the executable to a path, located in the PATH, so your OS can find it and makes it available as a command:

=== "Linux"
    ```sh
    ./install.sh
    ```
=== "Windows"
    ```sh
    .\install.bat
    ```

### Use
```sh
compose-generator [<command>]
```