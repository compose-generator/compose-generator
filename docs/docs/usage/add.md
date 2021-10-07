---
title: Add command
---

Sometimes you need to add one or more services to your Docker Compose configuration subsequently. For this you can use the `add` command of Compose Generator. It handles things like duplicate port assignments, service name collisions, Docker networking, etc. automatically for you.

## Usage
=== "Long form"
    Use the `add` command by executing:
    ```sh
    $ compose-generator add [service-name]
    ```
=== "Short form"
    Use the `add` command by executing:
    ```sh
    $ compose-generator a [service-name]
    ```

## Options
You can apply following options to the `add` command:

| Option       | Short | Description                                                               |
| ------------ | ----- | ------------------------------------------------------------------------- |
| `--advanced` | `-a`  | Enable advanced mode with advanced questions to allow more customization. |
| `--detached` | `-d`  | Run the compose configuration subsequently in detached mode.              |
| `--force`    | `-f`  | Skip safety checks and delete all files that may exist.                   |
| `--run`      | `-r`  | Run the compose configuration subsequently.                               |