# PascalCalc — Docker and API gateway

PascalCalc runs four containers: one Nginx container serving the built React frontend and acting as the API gateway, plus three independently deployed Go API containers. Each Go container runs its own executable and exposes only its own API endpoints on the Compose network.

## Prerequisites

Docker Desktop with Docker Compose v2. Run commands from the `calculator-app` directory.

## Start the application

```sh
docker compose up --build -d
docker compose ps
```

Open **http://localhost:8088**. To choose another host port, set `PASCALCALC_PORT=8090` before starting. The root `.env` file can retain `SERVER_PORT=8080`: Compose explicitly sets the correct service-specific ports (8081/8082/8083) inside each container. The root `.env` is **not** copied into the images.

## Verify through Nginx

```sh
curl http://localhost:8088/health
curl -X POST http://localhost:8088/api/v1/arithmetic/add -H 'Content-Type: application/json' -d '{"a":10,"b":20}'
curl -X POST http://localhost:8088/api/v1/math/power -H 'Content-Type: application/json' -d '{"a":2,"b":3}'
curl -X POST http://localhost:8088/api/v1/expressions/evaluate -H 'Content-Type: application/json' -d '{"expression":"10+20*2"}'
```

Expected results: gateway health OK, arithmetic result 30, math result 8, expression result 50.

## Routing

| Public route | Container |
| --- | --- |
| `/api/v1/arithmetic/*` | `arithmetic:8081` |
| `/api/v1/math/*` | `math:8082` |
| `/api/v1/expressions/*` | `expression:8083` |
| `/` | React static files served by Nginx |

The browser uses relative `/api` URLs. Nginx preserves the request path; the Go handlers do not need to change. The three Go services are reachable only on the internal Compose network, not directly from the host.

## Independent deployment demonstration

```sh
docker compose restart arithmetic
docker compose logs --tail=30 arithmetic
docker compose stop math
# Arithmetic and Expression endpoints still work; Math requests return 502.
docker compose start math
```

To rebuild just one Go service:

```sh
docker compose up -d --build --no-deps arithmetic
```

## Local development without Docker

Existing Vite development mode is unchanged. `npm run dev` still proxies `/api` to `BACKEND_API_URL` (default `http://localhost:8080`), the original combined Go server. The Docker deployment uses the Nginx gateway instead.

## Tests and cleanup

```sh
cd backend && go test ./... && cd ..
cd frontend && npm test -- --run && npm run build && cd ..
docker compose down
```

**Note:** The Docker image build and Compose runtime need to be verified on a machine with Docker installed; the packaging environment used to prepare this update does not have Docker.
