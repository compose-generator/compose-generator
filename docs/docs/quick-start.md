# Quick Start

Compose Generator offers several different methods to start with. Choose the one you need for your project below.

### Setup Compose Generator

??? summary "I want to use it for development"
    Compose Generator makes developing and especially testing much more seamless. Get started by choosing one of the following options:

    ??? summary "I have Docker already installed"
        Perfect! One thing less to care about ... Let's continue by installing Compose Generator:

        ??? summary "Install it on your development machine (recommended)"
            If you use Windows on your development machine, please refer to the guide on how to <a href="../install/windows" target="_blank">install Compose Generator on Windows</a> and if you work with Linux, you can visit the guide on how to  <a href="../install/linux" target="_blank">install Compose Generator on Linux</a>.

        ??? summary "Use it as Docker Container without installing"
            Please refer to the guide on how to <a href="../install/docker" target="_blank">use Compose Generator with Docker</a>.

        ??? summary "Install it on your development machine via NPM"
            If you have already installed NPM on your development machine, you can <a href="../install/npm" target="_blank">install Compose Generator via NPM</a>. The only thing NPM does, is to unpack the binary for your platform and architecture to the binaries directory of your system. Compose Generator will post-install predefined service templates and other essential files.

    ??? summary "Docker is not installed yet"
        No problem. Some commands of Compose Generator can even by used without having Docker installed, although it is recommended to have Docker installed.

        As Docker is required for most of the tasks for Compose Generator to cover, you should have Docker installed before using it. Compose Generator offers <a href="../usage/install" target="_blank">a simple command</a> to install Docker with Docker Compose.

        ??? summary "Install it on your development machine"
            If you use Windows on your development machine, please refer to the guide on how to <a href="../install/windows" target="_blank">install Compose Generator on Windows</a> and if you work with Linux, you can visit the guide on how to  <a href="../install/linux" target="_blank">install Compose Generator on Linux</a>.

        ??? summary "Install it on your development machine via NPM"
            If you have already installed NPM on your development machine, you can <a href="../install/npm" target="_blank">install Compose Generator via NPM</a>. The only thing NPM does, is to unpack the binary for your platform and architecture to the binaries directory of your system. Compose Generator will post-install predefined service templates and other essential files.

??? summary "I want to use it in production"
    Compose Generator can also be used to generate production ready Docker Compose configurations. To use it in a production environment we have two options: One-time use or more frequent usage.

    ??? summary "Use it as Docker Container without installing (recommended for one-time use)"
        Please refer to the guide on how to <a href="../install/docker" target="_blank">use Compose Generator with Docker</a>.

    ??? summary "Install it on your development machine (recommended when using more frequently)"
        Installing Compose Generator with the native package manager is recommended due to the enhanced ability to udate to newer versions and to integrate in potentially existing devops workflows.

        If you use Windows on your development machine, please refer to the guide on how to <a href="../install/windows" target="_blank">install Compose Generator on Windows</a> and if you work with Linux, you can visit the guide on how to  <a href="../install/linux" target="_blank">install Compose Generator on Linux</a>.

<!--??? summary "I want to use it for CI/CD"
    You want to use it for development and do not have Docker installed-->

### Generate your first Docker Compose Configuration

The first step is, to start your Docker instance on your host system. That given, follow one guide below to proceed:

??? question "I want to start from scratch"
    If you haven't started to work on your project you can simply run `$ compose-generator --with-instructions` to generate a compose project from scratch.
    
    Select if you want to only have a compose configuration for development or for development and production. If you choose the latter, Compose Generator creates a compose config with a reverse proxy and a TLS certificate helper service and provides two Docker profiles called `dev` and `prod` that you can run `$ docker compose -p dev up` or respectively `$ docker compose -p prod up`.

    After choosing the production-readiness, Compose Generator asks you for the services that you want to use. Please select one or more by navigating up and down and selecting by hitting the space key. Continue with the enter key. Compose-Generator will ask you some questions on how to configure the services, based on your selection.

    You will get asked for your selection of services in 4 or 6 categories, depending on your selection for the production-readiness. The categories will come up in following order: (`proxy`, `tls helper`,) `frontend`, `backend`, `database`, `db admin`.

    After you answered all questions, Compose Generator starts to generate the compose configuration for you and saves the provided outcome into the current directory. After finishing, you can run `$ docker compose up` to run Docker Compose with the generated configuration.

??? question "I already have project(s) to deploy by building the image(s) in-place (not recommended for production)"
    Clone the project to your system if not done already and execute `$ compose-generator --with-instructions` in the root directory of the project. Select `Custom service` plus the predefined templates for the companion services that you need.

??? question "I already have project(s) to deploy by pulling remote images"
    Clone the project to your system if not done already and execute `$ compose-generator --with-instructions` in the root directory of the project. Select `Custom service` plus the predefined templates for the companion services that you need.