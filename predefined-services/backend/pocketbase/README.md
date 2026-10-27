## PocketBase
PocketBase is an open source backend, based on the Go programming language. It consists of:

-   embedded database (SQLite) with realtime subscriptions
-   built-in files and users management
-   convenient Admin dashboard UI
-   and simple REST-ish API

PocketBase offers several sdk clients to interact with the backend.

### Setup
PocketBase is considered as backend service and can therefore be found in backends collection, when generating the compose configuration with Compose Generator.

After starting the generated Docker Compose configuration, the Admin UI Dashboard is exposed at the path `<scheme>://<host>:<port>/_/`.