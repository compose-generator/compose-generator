---
title: Template command
---

### Save
#### Description
Saves the compose configuration in the current directory.

!!! info
    Please note, that "templates" are something different than "predefined templates", which you can use with the [generate command](../generate). Thus, neither the list of predefined templates of the [generate command](../generate) contains any templates nor the list of templates contains any predefined templates.

#### Usage
=== "Long form"
    Use the `save` sub-command by executing:
    ```sh
    $ compose-generator template save [template-name]
    ```
=== "Short form"
    Use the `save` sub-command command by executing:
    ```sh
    $ compose-generator t s [template-name]
    ```

#### Options
You can apply following options to the `save` sub-command:

| Option    | Short | Description                                                 |
| --------- | ----- | ----------------------------------------------------------- |
| `--force` | `-f`  | Skip safety checks and overwrite all files, that may exist. |
| `--stash` | `-s`  | Remove configuration files after saving the template.       |

### Load
#### Description
Loads a compose configuration from a custom template.

#### Usage
=== "Long form"
    Use the `load` sub-command by executing:
    ```sh
    $ compose-generator template load [template-name]
    ```
=== "Short form"
    Use the `load` sub-command command by executing:
    ```sh
    $ compose-generator t l [template-name]
    ```

#### Options
You can apply following options to the `load` sub-command:

| Option    | Short | Description                                                 |
| --------- | ----- | ----------------------------------------------------------- |
| `--force` | `-f`  | Skip safety checks and overwrite all files, that may exist. |