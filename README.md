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
