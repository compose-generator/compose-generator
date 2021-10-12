---
title: Remove command
---

Compose Generator can help you to remove services from an existing Docker Compose configuration without leaving unused networks, service dependencies or other stuff, resulting in chaos. The `remove` command can help you to remove a single or multiple services at once from your configuration.

## Usage
=== "Long form"
    Use the `remove` command by executing:
    ```sh
    $ compose-generator remove [service-name]...
    ```
=== "Short form"
    Use the `remove` command by executing:
    ```sh
    $ compose-generator r [service-name]...
    ```

## Options
You can apply following options to the `remove` command:

| Option           | Short | Description                                                               |
| ---------------- | ----- | ------------------------------------------------------------------------- |
| `--advanced`     | `-a`  | Enable advanced mode with advanced questions to allow more customization. |
| `--detached`     | `-d`  | Run the compose configuration subsequently in detached mode.              |
| `--force`        | `-f`  | Skip safety checks and delete all files that may exist.                   |
| `--run`          | `-r`  | Run the compose configuration subsequently.                               |
| `--with-volumes` | `-v`  | Also remove all associated volumes of the stated services on disk.        |