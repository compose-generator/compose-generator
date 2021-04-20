# Quick Start

Compose Generator offers several different methods to use it and it depends on your context and intention, what we can recommend you to do.

### Setup Compose Generator

??? summary "I want to use it for development"
    Compose Generator makes developing and especially testing much more seamless. Get started by choosing of the following options:

    ??? summary "I have Docker already installed"
        Perfect! One thing less to care about ... Let's continue by installing Compose Generator:

        ??? summary "Use it as Docker Container without installing"
            Please refer to the guide on how to [use Compose Generator with Docker](../install/docker).

        ??? summary "Install it on your development machine (recommended)"
            If you use Windows on your development machine, please refer to the guide on how to [install Compose Generator on Windows](../install/windows) and if you work with Linux, you can visit the guide on how to [install Compose Generator on Linux](../install/linux).

        ??? summary "Install it on your development machine via NPM"
            If you have already installed NPM on your development machine, you can [install Compose Generator via NPM](../install/npm). The only thing NPM does, is to unpack the binary for your platform and architecture to the binaries directory of your system. Compose Generator will post-install predefined service templates and other essential files.

    ??? summary "I haven't got Docker yet"
        No problem. Some commands of Compose Generator can even by used without having Docker installed, although it is recommended to have Docker installed.

        ??? summary "Install it on your development machine"
            If you use Windows on your development machine, please refer to the guide on how to [install Compose Generator on Windows](../install/windows) and if you work with Linux, you can visit the guide on how to [install Compose Generator on Linux](../install/linux).

        ??? summary "Install it on your development machine via NPM"
            If you have already installed NPM on your development machine, you can [install Compose Generator via NPM](../install/npm). The only thing NPM does, is to unpack the binary for your platform and architecture to the binaries directory of your system. Compose Generator will post-install predefined service templates and other essential files.

        As Docker is required for most of the task, Compose Generator covers, you should have Docker installed before using it. Compose Generator offers [a simple command](../usage/install) to install both - Docker and Docker Compose - at once.

??? summary "I want to use it in production"
    Compose Generator can also be used to generate production ready Docker Compose configurations. To use it in a production environment we have two options: One-time use or more frequent usage.

    ??? summary "Use it as Docker Container without installing (recommended for one-time use)"
        Please refer to the guide on how to [use Compose Generator with Docker](../install/docker).

    ??? summary "Install it on your development machine (recommended when using more frequently)"
        Installing Compose Generator with the native package manager is recommended due to the enhanced ability to udate to newer versions and to integrate in potentially existing devops workflows.

        If you use Windows on your development machine, please refer to the guide on how to [install Compose Generator on Windows](../install/windows) and if you work with Linux, you can visit the guide on how to [insatll Compose Generator on Linux](../install/linux).

<!--??? summary "I want to use it for CI/CD"
    You want to use it for development and do not have Docker installed-->

### Generate your first Docker Compose Configuration

??? question "I want to start from scratch"
    To be extended ...

??? question "I already have project(s) to deploy by building images on the host system"
    To be extended ...

??? question "I already have project(s) to deploy by pulling remote images"
    To be extended ...