# NBA Reference API

This project provides a simple REST API backed by Postgres and written in Go. It currently exposes an endpoint to list NBA teams.

## Prerequisites

- **Go 1.24+** – see `go.mod` for the version used.
- **PostgreSQL 16** (available via the provided `docker-compose.yml`).
- Run the SQL migrations found in `database/migrations` before starting the API. You can run them with `make migrate-up` which expects `migrate` to be installed and `DB_URL` configured in the `Makefile`.

## Starting the API

1. Ensure Postgres is running (`docker-compose up -d db`).
2. Apply the migrations (`make migrate-up`).
3. Start the server:
   ```bash
   make run                # or
   go run cmd/api/main.go
   ```
   The server listens on port `8080`.

## Example

`GET /teams` returns all teams stored in the database.

Example request using `curl`:

```bash
curl http://localhost:8080/teams
```

Example response:

```json
[
  {
    "ID": "c2b92d2e-dc6f-4a4e-b19f-2a03067266f7",
    "TeamID": "atl",
    "Name": "Atlanta Hawks",
    "CreatedAt": "2024-07-05T12:00:00Z",
    "UpdatedAt": "2024-07-05T12:00:00Z"
  }
]
```
