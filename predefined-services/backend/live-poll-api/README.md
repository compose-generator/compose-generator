## Live-Poll API
Live-Poll is an open-source platform for taking surveys or polls. There is an [free hosted version online](https://www.live-poll.de) for the common use, but can also be self-hosted. This Compose Generator configuration is a good starting point for the usage of Live-Poll.

### Setup
Live-Poll comes as two parts: A frontend (categorized as a frontend service) and an API (this service).

The special thing when setting up Live-Poll API is, that it needs the login credentials to a SMTP server to send registration mails to the users.

Furthermore, the Live-Poll API requires a MySQL database to connect against. The data initialization happens automatically at the first startup. At the first startup, the Live-Poll API might restart a few times, because the database is not ready yet. This can take up to several minutes, depending on your system performance.