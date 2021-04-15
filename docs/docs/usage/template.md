---
title: Template command
---

You can use Compose Generator to save your Docker Compose configurations for later use and to restore them. Use the `template save` command to save your custom configuration and load it with `template load`.

!!! info
    Please note, that "templates" are something different than "predefined service templates", which you can use with the [generate command](../generate). Thus, neither the list of predefined service templates of the [generate command](../generate) contains any templates nor the list of templates contains any predefined service templates.

## Save template

Use the `template load` command to save custom configuration templates.

### Usage
=== "Long form"
    Use the `save` sub-command by executing:
    ```sh
    $ compose-generator template save [template-name]
    ```
=== "Short form"
    Use the `save` sub-command by executing:
    ```sh
    $ compose-generator t s [template-name]
    ```

### Options
You can apply following options to the `save` sub-command:

| Option              | Short | Description                                                 |
| ------------------- | ----- | ----------------------------------------------------------- |
| `--force`           | `-f`  | Skip safety checks and overwrite all files, that may exist. |
| `--stash`           | `-s`  | Remove configuration files after saving the template.       |
| `--with-dockerfile` | `-w`  | Save also Dockerfile in the template.                       |

## Load template

Use the `template load` command to load custom configuration templates again.

### Usage
=== "Long form"
    Use the `load` sub-command by executing:
    ```sh
    $ compose-generator template load [template-name]
    ```
=== "Short form"
    Use the `load` sub-command by executing:
    ```sh
    $ compose-generator t l [template-name]
    ```

### Options
You can apply following options to the `load` sub-command:

| Option              | Short | Description                                                            |
| ------------------- | ----- | ---------------------------------------------------------------------- |
| `--force`           | `-f`  | Skip safety checks and overwrite all files, that may exist.            |
| `--show`            | `-s`  | Do not load a template. Instead only list all templates and terminate. |
| `--with-dockerfile` | `-w`  | Load Dockerfile from template (if exists)                              |