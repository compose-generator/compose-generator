---
title: Install command
---

As Compose Generator acts as a helper for your Docker environment, it can also be used to install Docker and Docker Compose on your host machine with a single command. This helps to make your start much more seamless. With the lastest version of Docker, the Docker Compose cli is already bundled into the Docker installation and is used by Compose Generator by default.

!!! warning
    The install command does not work in combination with the dockerized version of Compose Generator. If you want to use the install command, please install Compose Generator on your host system. Please see the guides for [Linux](../../install/linux), [Windows](../../install/windows) or [NPM](../../install/npm) to do so.

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