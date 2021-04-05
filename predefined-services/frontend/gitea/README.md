## Gitea
Gitea is a tool for setting up own hosted Git repositories like GitHub or GitLab. It can be used in combination with PostgreSQL, MySQL, MariaDB, SQLite and MSSQL. For Compose Generator, we recommend the usage in combination with PostgreSQL, because the connection is tested the best.

### Setup
Gitea is considered as frontend service and can therefore be found in frontends collection, when generating the compose configuration with Compose Generator.

#### Generating a Gitea stack (can be skipped if already done)
Execute `$ compose-generator -ir` and answer all questions. Please select `Gitea` as frontend and one of `PostgreSQL`, `MySQL`, `MariaDB`, `SQLite`, `MSSQL` as database.

#### Install Gitea
After generating and starting the stack, you should be able to access Gitea via the port, you've set.

To install Gitea and connect it to your database, click on 'Register' in the top right corner. Here you can customize your Gitea installation. The following options are recommended:

Database type: Depends on the type, you've selected
Database credentials: Can be found in the `environment.env` file, that was generated.
Gitea HTTP Listen Port: Port you've set
Gitea Base URL: Host on which Gitea is accessible from outside

You also can add an admin user. All other settings can be left as default.