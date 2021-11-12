## Ghost
Ghost is an open-source blogging platform which can be self-hosted and is written in JavaScript. It is specialized on writing and publishing articles and blogposts.

### Setup
Ghost is considered as frontend service and can therefore be found in frontends collection, when generating the compose configuration with Compose Generator.

**Attention**: There is a bug in the Ghost MySQL connection flow, causing Ghost to not be able to connect to the database, complaining about the auth method. If this occurs to you, please select a MySQL version <=5.7.

After starting the Ghost Docker service, navigate to `<host>/ghost/` to get to the admin site. After creating an account there, you can start editing your articles.