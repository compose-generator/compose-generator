## Overleaf
Overleaf is a web-based editor to collaborate on LaTex projects.

### Setup
Overleaf is considered as frontend service and can therefore be found in frontends collection, when generating the compose configuration with Compose Generator.

### Create Admin user
To create the admin user, execute the following command on the Docker host system while the containers are running:
```sh
docker exec ${{PROJECT_NAME_CONTAINER}}-frontend-overleaf /bin/bash -c "cd /var/www/sharelatex; grunt user:create-admin --email=<admin-mail>"
```