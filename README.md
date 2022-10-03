## Migrations

### Generate migration

```bash
migrate create -ext sql -dir db/migrations -seq MIGRATION_NAME
```

### Migrate

```bash
migrate -path db/migrations -database "postgresql://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=disable" -verbose up
```

OR

```bash
bin/db_migrate
```

### Rollback

```bash
migrate -path db/migrations -database "postgresql://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=disable" -verbose down
```

OR

```bash
bin/db_rollback
```
