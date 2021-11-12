## Drupal
Drupal is a free and open-source content managment system (CMS) written in PHP. It can be used for blogs or personal websites.

### Setup
Ghost is considered as frontend service and can therefore be found in frontends collection, when generating the compose configuration with Compose Generator.

**Installation**: <br>
After launching the service, Drupal will prompt you with an installation wizard. On the step "Setup database", please select the database type (default is MySQL) and enter the database credentials, provided in `environment.env`. Don't forget to expand the advanced database configuration section to change the database host to `database-mysql` respecively `database-postgres` and optionally provide a table prefix. <br>
The subsequent installation process could take a while to finish.