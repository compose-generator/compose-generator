---
title: Generate command
---

### Description
You can use the `generate` command to generate Docker Compose configurations. You can generate your configurations from predefined stack templates or create your own stack.

!!! tip inline end
    You can later save / load your own templates with the [`template` command](../template).

### Usage
Use the `generate` command by executing `$ compose-generator` or `$ compose-generator generate`

### Options
You can apply following options to the command:

| Option        | Short | Description                                                                                                        |
| ------------- | ----- | ------------------------------------------------------------------------------------------------------------------ |
| `--advanced`  | `-a`  | Enabled advanced mode to show advanced questions to allow more customization                                       |
| `--demonized` | `-d`  | To run the compose configuration after generating in detached mode. For combined use with the `--run` / `-r` flag. |
| `--force`     | `-f`  | Skip safety checks and overwrite all files, that may exist.                                                        |
| `--run`       | `-r`  | To run the compose configuration after generating.                                                                 |