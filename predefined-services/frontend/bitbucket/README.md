## BitBucket
BitBucket is a source code management system for Git similar to GitHub. You can create repositories, set fine-granular permissions and work on your code with larger teams.

### Setup
BitBucket is considered as frontend service and can therefore be found in frontends collection, when generating the compose configuration with Compose Generator.

#### Generating a BitBucket stack (can be skipped if already done)
Execute `$ compose-generator -ir` and answer all questions. Please select `BitBucket` as frontend. BitBucket supports `PostgreSQL`, `MySQL`, `Oracle` or `MSSQL` as database, but Compose Generator currently supports only `PostgreSQL`.

#### Install BitBucket
After generating and starting the stack, you should be able to access BitBucket via the port, you've set.

To install BitBucket and connect it to your database, select 'extern'. The following options are recommended:

Database Type: Depends on the type, you've selected
Hostname: database-postgres
Port: leave default value
Database credentials: Can be found in the `environment.env` file, that was generated.