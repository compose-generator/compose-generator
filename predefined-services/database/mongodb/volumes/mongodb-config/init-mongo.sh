#!/bin/bash
mongo <<EOF
use $MONGO_INITDB_DATABASE
db.createUser({
    user: "$MONGODB_APPLICATION_USER",
    password: "$MONGODB_APPLICATION_USER_PW",
    roles: [{
        role: 'readWrite',
        db: '$MONGO_INITDB_DATABASE'
    }]
})
EOF