---
title: Install on Linux
---

### Install from repository
=== "Debian/Ubuntu"
    To install Compose Generator on Debian or Ubuntu, execute the following commands in your terminal:
    ```sh
    sudo apt-get install ca-certificates
    curl -fsSL https://server.chillibits.com/files/repo/gpg | sudo apt-key add -
	sudo add-apt-repository "deb https://repo.chillibits.com/$(lsb_release -is | awk '{print tolower($0)}')-$(lsb_release -cs) $(lsb_release -cs) main"
	sudo apt-get update
	sudo apt-get install compose-generator
    ```

=== "Fedora"
    To install Compose Generator on Fedora, execute the following commands in your terminal:
    ```sh
    sudo dnf -y install dnf-plugins-core
	sudo dnf config-manager --add-repo https://server.chillibits.com/files/repo/fedora.repo
	sudo dnf install compose-generator
    ```

=== "CentOS"
    To install Compose Generator on CentOS, execute the following commands in your terminal:
    ```sh
    sudo yum install -y yum-utils
	sudo yum-config-manager --add-repo https://server.chillibits.com/files/repo/centos.repo
	sudo yum install compose-generator
    ```

=== "Raspbian"
    To install Compose Generator on Raspbian, execute the following commands in your terminal:
    ```sh
    sudo apt-get install ca-certificates
    curl -fsSL https://server.chillibits.com/files/repo/gpg | sudo apt-key add -
	sudo echo "deb [arch=armhf] https://repo.chillibits.com/$(lsb_release -is | awk '{print tolower($0)}')-$(lsb_release -cs) $(lsb_release -cs) main" > /etc/apt/sources.list.d/chillibits.list
	sudo apt-get update
	sudo apt-get install compose-generator
    ```

    !!! warning
        The support for Raspbian is at the experimental stage. Please File a bug ticket on <a href="https://github.com/compose-generator/compose-generator/issues/new/choose" target="_blank">GitHub</a>

<!-- === "Alpine"
    To install Compose Generator on Alpine, execute the following commands in your terminal:
    ```sh
    apk update
    sh -c "echo 'https://repo.chillibits.com/alpine/$(cat \
        /etc/os-release | grep VERSION_ID | cut -d "=" -f2 | cut -d "." \
        -f1,2)/main'" >> /etc/apk/repositories
    wget -O /etc/apk/keys/chillibits.repo https://server.chillibits.com/files/repo/gpg
    apk add -y compose-generator
    ```

    !!! note
        If there occure any errors on the last step, please try the following instead
        ```sh
        apk add compose-generator --allow-untrusted
        ``` -->

### Install from package file
You also can install Compose Generator from a package file on your host system.

=== "Debian/Ubuntu/Raspbian"
    [Download amd64](https://github.com/compose-generator/compose-generator/releases/latest/download/compose-generator_amd64.deb){ .md-button .md-button--primary }
    [Download arm64](https://github.com/compose-generator/compose-generator/releases/latest/download/compose-generator_arm64.deb){ .md-button .md-button--primary }
    [Download armv5](https://github.com/compose-generator/compose-generator/releases/latest/download/compose-generator_armv5.deb){ .md-button .md-button--primary }
    [Download armv6](https://github.com/compose-generator/compose-generator/releases/latest/download/compose-generator_armv6.deb){ .md-button .md-button--primary }
    [Download armv7](https://github.com/compose-generator/compose-generator/releases/latest/download/compose-generator_armv7.deb){ .md-button .md-button--primary }

    To install it, execute the following command:
    ```sh
    dpkg -i <deb-file-name>
    ```

=== "Fedora/CentOS"
    [Download amd64](https://github.com/compose-generator/compose-generator/releases/latest/download/compose-generator_amd64.rpm){ .md-button .md-button--primary }
    [Download arm64](https://github.com/compose-generator/compose-generator/releases/latest/download/compose-generator_arm64.rpm){ .md-button .md-button--primary }
    [Download armv5](https://github.com/compose-generator/compose-generator/releases/latest/download/compose-generator_armv5.rpm){ .md-button .md-button--primary }
    [Download armv6](https://github.com/compose-generator/compose-generator/releases/latest/download/compose-generator_armv6.rpm){ .md-button .md-button--primary }
    [Download armv7](https://github.com/compose-generator/compose-generator/releases/latest/download/compose-generator_armv7.rpm){ .md-button .md-button--primary }

    To install it, execute the following command:
    ```sh
    rpm -U <rpm-file-name>
    ```

=== "Alpine"
    [Download amd64](https://github.com/compose-generator/compose-generator/releases/latest/download/compose-generator_amd64.apk){ .md-button .md-button--primary }
    [Download arm64](https://github.com/compose-generator/compose-generator/releases/latest/download/compose-generator_arm64.apk){ .md-button .md-button--primary }
    [Download armv5](https://github.com/compose-generator/compose-generator/releases/latest/download/compose-generator_armv5.apk){ .md-button .md-button--primary }
    [Download armv6](https://github.com/compose-generator/compose-generator/releases/latest/download/compose-generator_armv6.apk){ .md-button .md-button--primary }
    [Download armv7](https://github.com/compose-generator/compose-generator/releases/latest/download/compose-generator_armv7.apk){ .md-button .md-button--primary }

    To install it, execute the following command:
    ```sh
    apk add --allow-untrusted <apk-file-name>
    ```

### Use
```sh
compose-generator [<command>]
```