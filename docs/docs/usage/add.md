---
title: Add command
---

## Description
Adds a service to an existing compose file.

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

| Option       | Short | Description                                                     |
| ------------ | ----- | --------------------------------------------------------------- |
| `--advanced` | `-a`  | Show questions for advanced customization                       |
| `--detached` | `-d`  | To run the compose configuration after adding in detached mode. |
| `--force`    | `-f`  | Skip safety checks and delete all files that may exist.         |
| `--run`      | `-r`  | To run the compose configuration after adding.                  |