# MongoDB database
MongoDB is a common representative of NoSQL databases / document stores.

### Usage
Compose Generator will setup MongoDB so that an admin and an application database will be created on the first startup of the container. Furthermore an user for your application will be created and granted read/write access to the application database.

### Access manually
To access the data manually, you can use [MongoDBCompass](https://www.mongodb.com/products/compass). It provides you an installable tool to access your local and remote MongoDB instances.

### Known issues
On Windows, please use Powershell instead of the legacy Commmand Shell. Due to some reason, the MongoDB init script does not work, when the container was started from a Command Shell session.