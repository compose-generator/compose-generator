---
title: Install on Linux
---

### Install from repository
=== "Debian/Ubuntu"
    To install Compose Generator on Debian, execute the following commands in your terminal:
    ```sh
    sudo apt-get update
    sudo apt-get install -y apt-transport-https ca-certificates curl \
        gnupg-agent software-properties-common lsb-release
    curl -fsSL https://repo.chillibits.com/artifactory/debian/gpg | \
        sudo apt-key add -
    sudo add-apt-repository "deb https://repo.chillibits.com/artifactory/debian \
        $(lsb_release -cs) main"
    sudo sudo apt-get update
    sudo apt-get install -y compose-generator

    ```

=== "Fedora"
    To install Compose Generator on Fedora, execute the following commands in your terminal:
    ```sh
    sudo dnf -y install dnf-plugins-core
    sudo dnf config-manager --add-repo \
        https://repo.chillibits.com/artifactory/rpm/chillibits.repo
    sudo dnf -y install compose-generator

    ```

=== "CentOS"
    To install Compose Generator on CentOS, execute the following commands in your terminal:
    ```sh
    sudo yum install -y yum-utils
    sudo yum-config-manager --add-repo \
        https://repo.chillibits.com/artifactory/rpm/chillibits.repo
    sudo yum install -y compose-generator

    ```

=== "Alpine"
    To install Compose Generator on Alpine, execute the following commands in your terminal:
    ```sh
    apk update
    sh -c "echo 'https://repo.chillibits.com/artifactory/alpine/$(cat \
        /etc/os-release | grep VERSION_ID | cut -d "=" -f2 | cut -d "." \
        -f1,2)/main'" >> /etc/apk/repositories
    wget -O /etc/apk/keys/alpine.rsa.pub \
        https://repo.chillibits.com/artifactory/alpine/alpine.rsa.pub
    apk add -y compose-generator

    ```

    !!! note
        If there occure any errors on the last step, please try the following instead
        ```sh
        apk add compose-generator --allow-untrusted
        ```

=== "Raspbian"
    To install Compose Generator on Raspbian, execute the following commands in your terminal:
    ```sh
    sudo apt-get update
    sudo apt-get install -y apt-transport-https ca-certificates curl \
        gnupg-agent software-properties-common lsb-release
    curl -fsSL https://repo.chillibits.com/artifactory/debian/gpg | \
        sudo apt-key add -
    sudo add-apt-repository "deb https://repo.chillibits.com/artifactory/debian \
        $(lsb_release -cs) main"
    sudo sudo apt-get update
    sudo apt-get install -y compose-generator
    ```

### Install from package file
You also can install Compose Generator from a package file on your host system. You can download the package [here](https://github.com/compose-generator/compose-generator/releases/latest).

=== "Debian/Ubuntu"
    To install it, execute the following command:
    ```sh
    dpkg -i <deb-file-name>
    ```

=== "Fedora"
    To install it, execute the following command:
    ```sh
    rpm -U <rpm-file-name>
    ```

=== "CentOS"
    To install it, execute the following command:
    ```sh
    rpm -U <rpm-file-name>
    ```

=== "Alpine"
    To install it, execute the following command:
    ```sh
    apk add --allow-untrusted <apk-file-name>
    ```

=== "Raspbian"
    To install it, execute the following command:
    ```sh
    dpkg -i <deb-file-name>
    ```

### Use
```sh
compose-generator
```