## Migrations

### Generate migration

`migrate create -ext sql -dir db/migrations -seq MIGRATION_NAME`

### Migrate

`migrate -path db/migrations -database "postgresql://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=disable" -verbose up`

OR

`bin/db_migrate`

### Rollback

`migrate -path db/migrations -database "postgresql://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=disable" -verbose down`

OR

`bin/db_rollback`