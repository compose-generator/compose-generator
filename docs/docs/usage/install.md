---
title: Install command
---

## Description
Installs Docker and Docker compose to make your start more seamless.

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