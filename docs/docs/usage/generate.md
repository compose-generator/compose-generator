---
title: Generate command
---

## Description
You can use the `generate` command to generate Docker Compose configurations. The services are grouped by type (`frontend`, `backend`, `database`, `db-admin-tool`, `proxy`, `tls-helper`). You can select one ore more services of a certain type. 

!!! tip
    You can later save / load your own templates with the [`template` command](../template).

## Usage
=== "Root command"
    Use the `generate` command by executing:
    ```sh
    $ compose-generator
    ```
=== "Long form"
    Use the `generate` command by executing:
    ```sh
    $ compose-generator generate
    ```
=== "Short form"
    Use the `generate` command by executing:
    ```sh
    $ compose-generator g
    ```

## Options
You can apply following options to the `generate` command:

| Option                | Short       | Description                                                                   |
| --------------------- | ----------- | ----------------------------------------------------------------------------- |
| `--advanced`          | `-a`        | Enabled advanced mode to show advanced questions to allow more customization. |
| `--config <file>`     | `-c <file>` | Pass a configuration file with predefined answers. Works good for CI.[^1]     |
| `--detached`          | `-d`        | To run the compose configuration after generating in detached mode.           |
| `--force`             | `-f`        | Skip safety checks and overwrite all files, that may exist.                   |
| `--run`               | `-r`        | To run the compose configuration after generating.                            |
| `--with-instructions` | `-i`        | Generate a README.md file with usage instruction for predefined template.     |

[^1]:
    You can find an example configuration file [here](https://github.com/compose-generator/compose-generator/blob/main/media/example-config.yml).