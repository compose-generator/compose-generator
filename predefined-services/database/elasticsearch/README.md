## Elasticsearch


### Setup
#### Prerequisites
Elasticsearch needs some extra configuration on the Docker host system to grant more resources to the container. Please execute the following commands on your Docker host system: <br>

Linux:
```sh
sudo sysctl -w vm.max_map_count=262144
sudo sysctl -w fs.file-max=65536
ulimit -n 65536
ulimit -u 4096
```

Windows:
```sh
wsl -d docker-desktop
sysctl -w vm.max_map_count=262144
sysctl -w fs.file-max=65536
ulimit -n 65536
ulimit -u 4096
```