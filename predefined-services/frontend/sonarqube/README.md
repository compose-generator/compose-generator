## SonarQube
SonarQube is an open source product for continuous inspection of code quality.

### Setup
#### Prerequisites
SonarQube needs some extra configuration on the Docker host system to grant more resources to the container. Please execute the following commands on your Docker host system: <br>
```sh
sysctl -w vm.max_map_count=262144
sysctl -w fs.file-max=65536
ulimit -n 65536
ulimit -u 4096
```

#### Install SonarQube
After SonarQube is running, you can login with the default administrator credentials:

Username: `admin` <br>
Password: `admin`

