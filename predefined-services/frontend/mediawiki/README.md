## Mediawiki
Mediawiki is a free and open-source wiki software. Amongst others, it runs on Wikipedia, Wiktionary and Wikimedia Commons. Like WordPress, which is based on a similar licensing and architecture, it has become the dominant software in its category.

### Setup
Mediawiki is considered as frontend service and can therefore be found in frontend collection, when generating the compose configuration with Compose Generator.

For the default setup with MySQL, please enter following settings when configuring the database connection: <br>
**Database**: MariaDB, MySQL
**Server**: database-mysql
**Database name**: ${{MYSQL_DATABASE}}
**Database user**: ${{MYSQL_USER}}
**Database password**: Look up the generated password in environment.env

Mediawiki will serve you a `LocalSettings.php` configuration file. Copy this file to a new directory (e.g. `./volumes/mediawiki-settings/LocalSettings.php`). Additionally, you need to add a new volume to the `frontend-mediawiki` service in the compose file (`docker-compose.yml`):

```yml
- type: bind
  source: ./volumes/mediawiki-settings/LocalSettings.php
  target: /var/www/html/LocalSettings.php
```

After a restart the service should work fine.