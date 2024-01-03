# HETTA - Htmx + Echo + Templ + Tailwind + AlpineJS by Quteo

This is a example project for getting following technologies to work with golang.

- HTMX
- Echo
- Templ
- Tailwind
- AlpineJS
- PostgreSQL

Graciously developed by https://quteo.com

As there is a bunch of things that need to be generated and built for a full server restart, we use Air to handle all that for us.

## Prequisities

```shell
npm install -D tailwindcss
go install github.com/cosmtrek/air@latest
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Run

```shell
docker compose up -d   # for the DB
migrate -database "postgres://admin:admin@localhost:5432/postgres?sslmode=disable" -path db/migrations up
air server --port 3000
```
