# PascalCalc

**Classic inspiration. Modern calculation.**

PascalCalc is a browser-based calculator built with **React, TypeScript, and Vite**, backed by **Go REST APIs**. It provides a keypad-driven **Classic Mode** for evaluating expressions and a **Precision Mode** for selecting arithmetic and mathematical operations directly. The production-style Docker Compose deployment uses **Nginx** to serve the frontend and route requests to three independently deployable Go services.

## 1. Features

- **Classic Mode:** arithmetic expressions, parentheses, operator precedence, unary signs, exponentiation, and nested `sqrt(...)` expressions.
- **Precision Mode:** addition, subtraction, multiplication, division, power, square root, and percentage.
- JSON-based, versioned REST APIs with input validation and meaningful errors.
- Separate Arithmetic, Math, and Expression Go services.
- Centralized React API client and Nginx API gateway.
- Unit tests, coverage reports, and Docker-based deployment.
- **Expression limits:** Classic Mode input is capped at **256 JavaScript string characters**; the backend Expression Service independently defaults to **1024 UTF-8 bytes**. These are intentionally documented as separate constraints in the uploaded implementation.

## 2. Technology stack and architecture

| Layer | Technology | Responsibility |
| --- | --- | --- |
| Frontend | React, TypeScript, Vite | Classic and Precision calculator interfaces |
| API client | TypeScript | Centralized requests, API routes, response/error handling |
| Gateway | Nginx | Serves frontend and proxies `/api/v1/...` routes |
| Arithmetic | Go (`8081`) | Add, subtract, multiply, divide |
| Math | Go (`8082`) | Power, square root, percentage |
| Expression | Go (`8083`) | Tokenize, parse, and evaluate expressions |
| Deployment | Docker Compose | Runs gateway and three backend containers |
 AI-Assisted Development | ChatGPT (OpenAI) | Requirements analysis, brainstorming, architecture design, code generation, unit-test assistance, code review, and documentation |

**Docker request flow:** Browser → Nginx (`localhost:8088`) → appropriate Go microservice → JSON response.

The repository also includes a combined Go server (`backend/cmd/api`) for convenient local development on port `8080`. Docker Compose deploys the three standalone services instead.

## 3. Prerequisites

For Docker-based use:
- Docker Desktop with Docker Compose v2.

For local development:
- Go toolchain compatible with `backend/go.mod`.
- Node.js compatible with the installed Vite version (Node.js 22+ recommended) and npm.

## 4. Quick start — Docker Hub images (ARM64)

The prebuilt Docker Hub deployment starts PascalCalc **without compiling Go or React**. The published frontend image currently supports **linux/arm64 only** (for example, Apple Silicon Macs). The three backend images support both `linux/amd64` and `linux/arm64`. On an AMD64 machine, use the [build-from-source option](#5-build-from-source-with-docker) until an AMD64 frontend image is published.

**Prerequisites:** Docker with Compose v2 and an ARM64 Docker environment.

```bash
git clone https://github.com/sandipa1982/pascalcalc.git
cd pascalcalc
docker compose -f compose.hub.yaml pull
docker compose -f compose.hub.yaml up -d
docker compose -f compose.hub.yaml ps
```

Open **http://localhost:8088**. Verify the gateway:

```bash
curl http://localhost:8088/health
# {"status":"ok","component":"gateway"}
```

To run alongside another PascalCalc deployment using port 8088, use port 8089 instead:

```bash
PASCALCALC_PORT=8089 docker compose -f compose.hub.yaml up -d
curl http://localhost:8089/health
```

**Published Docker Hub images:**

| Component | Image | Published platforms |
| --- | --- | --- |
| Arithmetic | `sandipain/pascalcalc-arithmetic:1.0` | AMD64, ARM64 |
| Math | `sandipain/pascalcalc-math:1.0` | AMD64, ARM64 |
| Expression | `sandipain/pascalcalc-expression:1.0` | AMD64, ARM64 |
| Frontend + Nginx | `sandipain/pascalcalc-web:1.0-arm64` | ARM64 |

The published frontend image includes the Nginx gateway configuration; the Compose network uses the service names `arithmetic`, `math`, and `expression`. Backend ports remain internal.

Useful commands:

```bash
docker compose -f compose.hub.yaml logs -f
docker compose -f compose.hub.yaml ps
docker compose -f compose.hub.yaml down
```

**Verification on 19 October 2026:** Four published-image containers were running, the gateway returned `{"status":"ok","component":"gateway"}` on port 8089, and browser calculations worked. Service-specific `/api/v1/math/health` and `/api/v1/expressions/health` returned 404; those URLs are not the documented calculation endpoints. Use the POST examples below to test the calculation APIs.

> **Repository requirement:** Commit `compose.hub.yaml` alongside this README; it is needed for the Docker Hub commands above.

## 5. Build from source with Docker

This option builds all four images locally and can be used on AMD64 or ARM64 hosts.

From the repository root:

```bash
docker compose up --build -d
docker compose ps
curl http://localhost:8088/health
```

Open **http://localhost:8088** in a browser.

Stop the containers:

```bash
docker compose down
```

The gateway is published on port `8088` by default. To change the host port:

```bash
PASCALCALC_PORT=8090 docker compose up --build -d
```

The Go services communicate over the internal Compose network and are not individually published to the host. See [Docker setup](DOCKER_SETUP.md).

## 6. Local development

### Option A — Combined Go API and Vite

In terminal 1, from the repository root:

```bash
cd backend
go run ./cmd/api
```

By default, the combined API listens on **http://localhost:8080**. A root `.env` file or environment variables can override this setting.

In terminal 2, from the repository root:

```bash
cd frontend
npm ci
npm run dev
```

Open **http://localhost:5173**. Vite proxies `/api` requests to `BACKEND_API_URL` (default `http://localhost:8080`).

### Option B — Independently running Go services

From `backend/`, start each service in a separate terminal:

```bash
SERVER_PORT=8081 go run ./cmd/arithmetic-api
SERVER_PORT=8082 go run ./cmd/math-api
SERVER_PORT=8083 go run ./cmd/expression-api
```

These are **three separate terminal commands**, not a sequential single-terminal script. For a single browser-facing API origin across all services, use the Docker Nginx gateway. The Vite development proxy targets one backend address, so the combined API (Option A) is simplest for local development.

Copy `.env.example` to `.env` if you need to customize settings. Do not commit secrets or local-only environment values.

## 7. REST API and examples

The Docker gateway base URL is **http://localhost:8088**. Calculation routes accept `POST` with JSON. Successful arithmetic/math requests return `{"result":number}`; expression evaluation returns `{"expression":"...","result":number}`.

| Service | Endpoint | Example request | Example response (HTTP 200) |
| --- | --- | --- | --- |
| Arithmetic | `POST /api/v1/arithmetic/add` | `{"a":10,"b":20}` | `{"result":30}` |
| Arithmetic | `POST /api/v1/arithmetic/subtract` | `{"a":20,"b":5}` | `{"result":15}` |
| Arithmetic | `POST /api/v1/arithmetic/multiply` | `{"a":6,"b":7}` | `{"result":42}` |
| Arithmetic | `POST /api/v1/arithmetic/divide` | `{"a":20,"b":4}` | `{"result":5}` |
| Math | `POST /api/v1/math/power` | `{"a":2,"b":3}` | `{"result":8}` |
| Math | `POST /api/v1/math/sqrt` | `{"a":16}` | `{"result":4}` |
| Math | `POST /api/v1/math/percentage` | `{"a":20,"b":150}` | `{"result":30}` |
| Expression | `POST /api/v1/expressions/evaluate` | `{"expression":"2+4+sqrt(16)"}` | `{"expression":"2+4+sqrt(16)","result":10}` |

**Arithmetic example**

```bash
curl -X POST http://localhost:8088/api/v1/arithmetic/add \
  -H 'Content-Type: application/json' \
  -d '{"a":10,"b":20}'
```

**Math example**

```bash
curl -X POST http://localhost:8088/api/v1/math/sqrt \
  -H 'Content-Type: application/json' \
  -d '{"a":16}'
```

**Expression example**

```bash
curl -X POST http://localhost:8088/api/v1/expressions/evaluate \
  -H 'Content-Type: application/json' \
  -d '{"expression":"2+4+sqrt(16)"}'
```

**Error example — division by zero**

```bash
curl -i -X POST http://localhost:8088/api/v1/arithmetic/divide \
  -H 'Content-Type: application/json' \
  -d '{"a":10,"b":0}'
```

The API returns HTTP `400` with a JSON `error` field for invalid calculation input; unsupported calculation methods return HTTP `405`. Health endpoints use `GET /health`.

For all routes, required fields, validation rules, request/response payloads, and error cases, see [API Documentation](docs/API_DOCUMENTATION.md).

## 8. Configuration

PascalCalc supports three execution environments:

1. **Docker Hub deployment:** Runs prebuilt images using `compose.hub.yaml`.
2. **Docker build from source:** Builds the application locally using `compose.yaml`.
3. **Local development:** Runs Go services and the React/Vite frontend directly.

Settings are supplied through environment variables, Docker Compose files, `.env.example`, and application defaults.

### 8.1 Configuration variables

| Variable | Default | Purpose |
| --- | --- | --- |
| `PASCALCALC_PORT` | `8088` | Host port for Nginx in both Docker Compose deployments |
| `SERVER_HOST` | `0.0.0.0` in Docker | Go HTTP listener address |
| `SERVER_PORT` | `8080` combined; `8081`/`8082`/`8083` in Compose | Go HTTP listener port |
| `MAX_REQUEST_BYTES` | `4096` | Maximum JSON request body size |
| `MAX_EXPRESSION_LENGTH` | `1024` | Backend expression limit in UTF-8 bytes |
| `HTTP_READ_TIMEOUT` | `5s` | Backend HTTP read timeout |
| `HTTP_WRITE_TIMEOUT` | `10s` | Backend HTTP write timeout |
| `HTTP_IDLE_TIMEOUT` | `60s` | Backend HTTP idle timeout |
| `BACKEND_API_URL` | `http://localhost:8080` | Vite development proxy target |
| `FRONTEND_PORT` | `5173` | Vite development server port |
| `VITE_API_BASE_URL` | `/api` | Frontend API base path (build-time setting) |
| `VITE_API_TIMEOUT_MS` | `10000` | Frontend API timeout setting |

### 8.2 Docker Hub configuration

The published Docker images use runtime environment variables supplied by `compose.hub.yaml`. For example:

```bash
PASCALCALC_PORT=8089 MAX_REQUEST_BYTES=8192 \
  docker compose -f compose.hub.yaml up -d
```

Open `http://localhost:8089` to use the application.

The Go services receive their configuration at container startup. The React frontend is already compiled into the published image; changing Vite build-time settings requires rebuilding and republishing the frontend image.

### 8.3 Docker build-from-source configuration

Build and start all services locally:

```bash
docker compose up --build -d
```

Override the published gateway port:

```bash
PASCALCALC_PORT=8090 docker compose up --build -d
```

### 8.4 Local development configuration

The Vite development server proxies `/api` requests to `BACKEND_API_URL` (default `http://localhost:8080`). Configure local values using the appropriate environment file or shell variables. See [Local development](#6-local-development).

### 8.5 Configuration considerations

- Compose-supplied backend environment variables take precedence over backend defaults.
- Backend services communicate using their Compose service names over the internal Docker network.
- Only the Nginx gateway is published to the host by default.
- Classic Mode limits expression input to 256 JavaScript string characters; the Expression Service independently defaults to 1024 UTF-8 bytes.
- The published frontend image currently supports ARM64; backend images support AMD64 and ARM64.

## 9. Testing and code coverage

### Go backend

```bash
cd backend
go test ./... -count=1 -coverprofile=coverage.out
go tool cover -func=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

### React frontend

```bash
cd frontend
npm ci
npm run test -- --run
npm run test -- --run --coverage
npm run build
```

The repository declares `@vitest/coverage-v8` as a frontend development dependency.

**Previously reported local results (19 October 2026):**

| Metric | Result |
| --- | ---: |
| Go backend statement coverage | **80.4%** |
| React test files | **2 passed** |
| React unit tests | **14 passed / 14 total** |
| React statement coverage | **84.44%** |
| React branch coverage | **74.13%** |
| React function coverage | **86.95%** |
| React line coverage | **85.93%** |

These values are from the developer's local command output, not from tests rerun during README preparation. Regenerate coverage after any code changes.

## 10. Key design decisions and trade-offs

1. **Three services by capability:** Arithmetic, Math, and Expression are independently deployable and reusable across both UI modes. This adds infrastructure complexity compared with a monolith but demonstrates clear service boundaries.
2. **One Go module with separate entry points:** Shared interfaces and code reduce duplication, while standalone binaries permit independent container deployment. Shared-module changes may still affect multiple services.
3. **Nginx gateway:** Presents a single browser origin, centralizes routing, and avoids exposing internal Go service ports in Docker Compose.
4. **Centralized frontend API client:** Keeps transport details and endpoint configuration outside calculator UI components.
5. **Parser-based expression evaluation:** Tokenization and parsing support precedence, parentheses, unary operators, and functions without evaluating arbitrary source code.
6. **Layered validation:** UI input checks improve usability, while Go handlers and calculation services independently validate requests and numeric operations.
7. **Stateless JSON APIs:** No database is required for calculator operations. Go `float64` results have normal floating-point limitations; this is not an arbitrary-precision decimal engine.
8. **Testing and documentation first:** Unit tests, coverage, explicit API contracts, and phase-wise AI prompts support 
review and reproducibility.
9. **Numeric input precision and 15-digit limit:** PascalCalc adopts a maximum of 15 digits for numeric operands as a product-level constraint. This conservative limit is intended to reduce precision-related surprises when using Go's float64 arithmetic and keep calculator inputs manageable. It does not guarantee exact decimal representation or arithmetic. Arbitrary-precision decimal calculations are outside the current scope.

10. **Expression-length limits and resource protection:** Classic Mode limits expression input to 256 JavaScript string characters to maintain usability and prevent excessively long expressions. The backend Expression Service independently applies a configurable limit of 1024 UTF-8 bytes by default to bound request processing. These separate validation layers provide defense in depth, with the trade-off that longer expressions may be rejected even when mathematically valid.

## 11. Repository structure

```text
calculator-app/
├── backend/
│   ├── cmd/                     # Combined and standalone service entry points
│   └── internal/                # Arithmetic, math, expression, API, config, server
├── frontend/
│   └── src/                     # React modes, API client, configuration, tests
├── deploy/nginx/               # Nginx gateway configuration
├── docs/
│   ├── prompts/
│   │   ├── Prompts-Coding-backend/
│   │   ├── Prompts-Coding-frontend/
│   │   ├── brainstorming_prompts.md
│   │   └── architecture_design_prompts.md
│   ├── SOLUTION_APPROACH.md
│   ├── overall_architecture.md
│   ├── backend-hld.md
│   ├── API_DOCUMENTATION.md
│   ├── FRONTEND_BLUEPRINT.md
│   └── FRONTEND_DESIGN.md
├── compose.yaml
├── compose.hub.yaml           # Published-image deployment (ARM64 frontend)
├── .env.example
├── DOCKER_SETUP.md
└── README.md
```

## 12. Design documents and AI prompts — phase-wise

The documentation follows the project's input → design → implementation progression. Links below reflect the **actual lowercase `docs/` directory and nested `docs/prompts/` paths in the uploaded project**, not VS Code-local URLs.

| Phase | Document / prompts | Role in development |
| --- | --- | --- |
| **1. Brainstorming** | [Brainstorming Prompts](docs/prompts/brainstorming_prompts.md) → [Solution Approach](docs/SOLUTION_APPROACH.md) | Explore alternatives and produce the proposed solution approach |
| **2. System Architecture** | [Architecture Design Prompts](docs/prompts/architecture_design_prompts.md) → [Overall Architecture](docs/overall_architecture.md) | Use the Solution Approach as input to architecture decisions |
| **3. Backend Design** | [Backend HLD](docs/backend-hld.md) | Define Go services, responsibilities, and contracts |
| **4. Backend Implementation** | [Backend Coding Prompts](docs/prompts/Prompts-Coding-backend/) | Project foundation, arithmetic, math, expression, HTTP API, infrastructure, and entry points |
| **5. API Documentation** | [API Documentation](docs/API_DOCUMENTATION.md) | Record implemented routes, request/response JSON, errors, and examples |
| **6. Frontend Design** | [Frontend Blueprint](docs/FRONTEND_BLUEPRINT.md) → [Frontend Design](docs/FRONTEND_DESIGN.md) | Define UI modes, components, and API integration |
| **7. Frontend Implementation** | [Frontend Coding Prompts](docs/prompts/Prompts-Coding-frontend/) | Project foundation, API configuration/client, calculators, app composition, and styles |
| **8. Deployment and Verification** | [Docker Setup](DOCKER_SETUP.md) and [Testing instructions](#9-testing-and-code-coverage) | Deploy the services, verify routes, run tests, and reproduce coverage |

### Backend coding prompt sequence

1. [Project Foundation](docs/prompts/Prompts-Coding-backend/01_Project_Foundation.md)
2. [Arithmetic](docs/prompts/Prompts-Coding-backend/02_Arithmetic.md)
3. [Math](docs/prompts/Prompts-Coding-backend/03_Math.md)
4. [Expression](docs/prompts/Prompts-Coding-backend/04_Expression.md)
5. [HTTP API](docs/prompts/Prompts-Coding-backend/05_HTTP_API.md)
6. [Service Infrastructure](docs/prompts/Prompts-Coding-backend/06_Service_Infrastructure.md)
7. [Service Entry Points](docs/prompts/Prompts-Coding-backend/07_Service_Entry_Points.md)

### Frontend coding prompt sequence

1. [Project Foundation](docs/prompts/Prompts-Coding-frontend/01_Project_Foundation.md)
2. [API Configuration](docs/prompts/Prompts-Coding-frontend/02_API_Configuration.md)
3. [Centralized API Client](docs/prompts/Prompts-Coding-frontend/03_Centralized_API_Client.md)
4. [Classic Calculator](docs/prompts/Prompts-Coding-frontend/04_Classic_Calculator.md)
5. [Precision Calculator](docs/prompts/Prompts-Coding-frontend/05_Precision_Calculator.md)
6. [App Composition and Styles](docs/prompts/Prompts-Coding-frontend/06_App_Composition_and_Styles.md)

The prompt files document AI-assisted work and may include reconstructed prompts. Their presence does not, by itself, prove that a specific prompt was executed or that generated output was independently validated. The source code and tests are the implementation evidence.


