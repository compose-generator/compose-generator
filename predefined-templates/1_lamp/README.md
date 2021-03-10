# LAMP (Linux, Apache, MySQL, PhpMyAdmin)
This is the manual document for the `1_lamp` stack.

## Description
This stack deploys a website with PHP support, using an Apache2 webserver. This website has access to a MySQL database, which runs on the host called `database`. Furtermore, you also have access to the database by using PhpMyAdmin.

## Usage instructions
This is a stand-alone template. You only have to start it by executing `docker-compose up -d`.

## Volumes
-   `apache-logs`: This is where the access & error logs of Apache go
-   `apache-php-config`: Here you can put your custom php.ini configuration in
-   `mysql-data`: Data directory for MySQL. You can copy / backup this directory.
-   `mysql-logs`: This is where the logs of MySQL go

## Credits
Template created & maintained by @marcauberer