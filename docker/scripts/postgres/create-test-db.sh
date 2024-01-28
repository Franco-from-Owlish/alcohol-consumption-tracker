#!/bin/bash

set -e;
echo "Creating test database"
psql -U "${POSTGRES_USER}" << EOF
  CREATE DATABASE "${DB_NAME_TEST}";
  GRANT ALL PRIVILEGES ON DATABASE "${DB_NAME_TEST}" TO ${POSTGRES_USER};
EOF
echo "Created test database"