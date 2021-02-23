---
title: Remove command
---

### Description
Removed a service from an existing compose file. You also can remove multiple services at once.

### Usage
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

### Options
You can apply following options to the `remove` command:

| Option           | Short | Description                                                                                                        |
| ---------------- | ----- | ------------------------------------------------------------------------------------------------------------------ |
| `--demonized`    | `-d`  | To run the compose configuration after generating in detached mode. For combined use with the `--run` / `-r` flag. |
| `--force`        | `-f`  | Skip safety checks and overwrite all files, that may exist.                                                        |
| `--run`          | `-r`  | To run the compose configuration after generating.                                                                 |
| `--with-volumes` | `-v`  | Also remove all associated volumes of the stated services.                                                         |