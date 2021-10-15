## ScyllaDB
ScyllaDB is a distributed database management system, which is compatible with but faster than Apache Cassandra. ScyllaDB exists as an enterprise and a community version.

## Setup
ScyllaDB is considered as database service and can therefore be found in database collection, when generating the compose configuration with Compose Generator.

### Create roles for role based access control (RBAC)
Compose Generator enables password protection per default, but you have to create the user credentials yourself. Use the [RBAC guide](https://docs.scylladb.com/operating-scylla/security/rbac-usecase) for doing that.