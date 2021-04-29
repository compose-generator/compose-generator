# Supported services

The following table shows all services, which Compose Generator currently supports out of the box as predefined service template. To use these templates, please refer to the [generate command](../usage/generate).

| Frontend           | Backend              | Database           | Database Admin     |
| ------------------ | -------------------- | ------------------ | ------------------ |
| `Angular`          | `Flask`              | `Cassandra`        | `Adminer`          |
| `BitBucket`        | `Gin Gonic`          | `Elasticsearch`    | `Elasticsearch HQ` |
| `Common Website`   | `Minecraft Server`   | `FaunaDB`          | `pgAdmin`          |
| `Gitea`            | `PHP`                | `InfluxDB`         | `PhpMyAdmin`       |
| `Jenkins`          | `Rocket`             | `MariaDB`          | to be extended ... |
| `Jira`             | `Spring with Gradle` | `MongoDB`          |                    |
| `Nextcloud`        | `Spring with Maven`  | `MySQL`            |                    |
| `Owncloud`         | to be extended ...   | `Neo4j`            |                    |
| `React`            |                      | `OrientDB`         |                    |
| `SonarQube`        |                      | `PostgreSQL`       |                    |
| `Vue`              |                      | to be extended ... |                    |
| `Wordpress`        |                      |                    |                    |
| `YouTrack`         |                      |                    |                    |
| to be extended ... |                      |                    |                    |

If you miss a predefined template and you want to create one for the public, please first read our [contribution guidelines](../contributing) and then continue reading the instructions on how to [contribute a service template](https://github.com/compose-generator/compose-generator/blob/docs/supported-services-page/predefined-services/README.md). Fork the repository, create the template and open a pr to the `dev` branch. The community is thankful for every predefined template!