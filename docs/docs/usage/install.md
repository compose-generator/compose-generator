---
title: Install command
---

As Compose Generator acts as helper for your Docker environment, it also can be used to install Docker and Docker Compose on your host machine with a single command. This helps to make your start much more seamless.

!!! warning
    The install commands does not work in combination with the dockerized version of Compose Generator. If you want to use the install command, please install Compose Generator on your host system. Please see the guides for [Linux](../../install/linux), [Windows](../../install/windows) or [NPM](../../install/npm) to do so.

## Usage
=== "Long form"
    Use the `install` command by executing:
    ```sh
    $ compose-generator install
    ```
=== "Short form"
    Use the `install` command by executing:
    ```sh
    $ compose-generator i
    ```

## Options
You can apply following options to the `install` command:

| Option             | Short | Description                     |
| ------------------ | ----- | ------------------------------- |
| `--only-compose`   | `-c`  | Only install Docker Compose[^1] |
| `--only-docker`    | `-d`  | Only install Docker[^1]         |

[^1]:
    Only works on Linux / MacOS. Docker Desktop for Windows already comes with builtin support for Docker Compose.