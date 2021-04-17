## Jira
Jira is a tool for project management by Atlassian. It can be used by agile teams to track their tasks and activities.

### Setup
Jira is considered as frontend service and can therefore be found in frontends collection, when generating the compose configuration with Compose Generator.

#### Generating a Gitea stack (can be skipped if already done)
Execute `$ compose-generator -ir` and answer all questions. Please select `Gitea` as frontend and one of `PostgreSQL`, `MySQL`, `Oracle` or `MSSQL` as database.

#### Install Jira
After generating and starting the stack, you should be able to access Jira via the port, you've set.

To install Jira and connect it to your database, click on 'I'll set it up myself' and select 'My Own Database'. The following options are recommended:

Database Type: Depends on the type, you've selected#
Hostname: database-postgres / database-mysql / database-oracle / database-mssql
Database credentials: Can be found in the `environment.env` file, that was generated.