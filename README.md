### Local development

Local development uses Docker Postgres.

Default local settings:

- Postgres version: 16
- Container name: servicemaster-postgres
- Host: localhost
- Port: 5432
- Database: servicemaster_dev
- User: servicemaster
- Password: servicemaster

Start local Postgres:

```sh
docker run --name servicemaster-postgres \
  -e POSTGRES_USER=servicemaster \
  -e POSTGRES_PASSWORD=servicemaster \
  -e POSTGRES_DB=servicemaster_dev \
  -p 5432:5432 \
  -d postgres:16
```

## Database migrations and queries

ServiceMaster uses Goose for SQL migrations and sqlc for typed Go query generation.

Directories:

- `db/migrations`: Goose migration files
- `db/queries`: handwritten SQL queries for sqlc
- `internal/store`: generated sqlc Go code

Apply local migrations:

```sh
goose -dir db/migrations postgres "$DATABASE_URL" up
```
