#!/bin/sh
set -a
. ../../.env
set +a
docker exec -i standard_app_dev-db-1 psql -U $POSTGRES_USER -d $POSTGRES_DB < $1.sql
