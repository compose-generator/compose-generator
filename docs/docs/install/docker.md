---
title: Use with Docker
---

### Download
You don't have to pull the image first. You also can skip this step.
=== "Docker Hub"
    ```sh
    docker pull chillibits/compose-generator
    ```
=== "GitHub Container Registry"
    ```sh
    docker pull ghcr.io/chillibits/compose-generator
    ```

### Use
=== "Linux - Docker Hub"
    ```sh
    docker run --rm -it -v /var/run/docker.sock:/var/run/docker.sock -v $(pwd):/cg/out chillibits/compose-generator [<command>]
    ```
=== "Linux - GHCR"
    ```sh
    docker run --rm -it -v /var/run/docker.sock:/var/run/docker.sock -v $(pwd):/cg/out ghcr.io/chillibits/compose-generator [<command>]
    ```
=== "Windows - Docker Hub"
    ```sh
    docker run --rm -it -v /var/run/docker.sock:/var/run/docker.sock -v ${pwd}:/cg/out chillibits/compose-generator [<command>]
    ```
=== "Windows - GHCR"
    ```sh
    docker run --rm -it -v /var/run/docker.sock:/var/run/docker.sock -v ${pwd}:/cg/out ghcr.io/chillibits/compose-generator [<command>]
    ```

### Customize
#### Custom output path
You can use another output path by replacing `${pwd}` with a custom path.

!!! example
    ```sh
    docker run --rm -it -v ./project:/cg/out chillibits/compose-generator
    ```

#### Expose template directory
You can pass another volume to save your custom templates to a directory on the host system.

!!! example
    ```sh
    docker run --rm -it -v ${pwd}:/cg/out -v ~/cg-templates:/cg/templates chillibits/compose-generator
    ```