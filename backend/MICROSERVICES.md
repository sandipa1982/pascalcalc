# PascalCalc — independent Go services

Run each command in a separate terminal from `backend/`:

```bash
go run ./cmd/arithmetic-api  # default port 8081
go run ./cmd/math-api        # default port 8082
go run ./cmd/expression-api  # default port 8083
```

All three use the same Go module but run as separate OS processes, each exposing only its own API and `GET /health`.

| Service | Public API paths |
|---|---|
| Arithmetic | `POST /api/v1/arithmetic/add`, `/subtract`, `/multiply`, `/divide` |
| Math | `POST /api/v1/math/power`, `/sqrt`, `/percentage` |
| Expression | `POST /api/v1/expressions/evaluate` |

Quick checks:

```bash
curl -s http://localhost:8081/health
curl -s http://localhost:8082/health
curl -s http://localhost:8083/health
curl -s -X POST http://localhost:8081/api/v1/arithmetic/add -H 'Content-Type: application/json' -d '{"a":10,"b":20}'
curl -s -X POST http://localhost:8082/api/v1/math/power -H 'Content-Type: application/json' -d '{"a":2,"b":3}'
curl -s -X POST http://localhost:8083/api/v1/expressions/evaluate -H 'Content-Type: application/json' -d '{"expression":"2+3"}'
```

Run the test suite:

```bash
go test ./...
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
```

Configuration: `SERVER_HOST` (default `0.0.0.0`), `SERVER_PORT` (optional override), request limits and HTTP timeouts are shared through `internal/config`. **Important:** if `.env` or your shell sets `SERVER_PORT=8080`, it overrides the per-service defaults; unset it when running the three services locally, or set it separately in each terminal.

The original combined server remains available through `go run ./cmd/api` on its original default port 8080.

Frontend routing via Nginx and Docker Compose is planned for Step 7.5; the existing frontend is not yet switched to these three independent service ports.
