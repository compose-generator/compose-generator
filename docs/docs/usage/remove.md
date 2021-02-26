---
title: Remove command
---

## Description
Removes a service from an existing compose file. You also can remove multiple services at once.

## Usage
=== "Long form"
    Use the `remove` command by executing:
    ```sh
    $ compose-generator remove [service-name]
    ```
=== "Short form"
    Use the `remove` command by executing:
    ```sh
    $ compose-generator r [service-name]
    ```

## Options
You can apply following options to the `remove` command:

| Option           | Short | Description                                                                                                      |
| ---------------- | ----- | ---------------------------------------------------------------------------------------------------------------- |
| `--advanced`     | `-a`  | Show questions for advanced customization                                                                        |
| `--demonized`    | `-d`  | To run the compose configuration after removing in detached mode. For combined use with the `--run` / `-r` flag. |
| `--force`        | `-f`  | Skip safety checks and delete all files that may exist.                                                          |
| `--run`          | `-r`  | To run the compose configuration after removing.                                                                 |
| `--with-volumes` | `-v`  | Also remove all associated volumes of the stated services on disk.                                               |