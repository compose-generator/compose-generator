#!/bin/sh
set -e

mongo <<EOF
use $MONGO_INITDB_DATABASE
db.createUser({
    user: "$MONGODB_APPLICATION_USER",
    pwd: "$MONGODB_APPLICATION_USER_PW",
    roles: [{
        role: 'readWrite',
        db: '$MONGO_INITDB_DATABASE'
    }]
})
EOF