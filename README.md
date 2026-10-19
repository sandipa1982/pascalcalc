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

**Docker request flow:** Browser → Nginx (`localhost:8088`) → appropriate Go microservice → JSON response.

The repository also includes a combined Go server (`backend/cmd/api`) for convenient local development on port `8080`. Docker Compose deploys the three standalone services instead.

## 3. Prerequisites

For Docker-based use:
- Docker Desktop with Docker Compose v2.

For local development:
- Go toolchain compatible with `backend/go.mod`.
- Node.js compatible with the installed Vite version (Node.js 22+ recommended) and npm.

## 4. Quick start — Docker (recommended)

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

## 5. Local development

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

## 6. REST API and examples

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

## 7. Configuration

Key settings defined by the Go configuration, `.env.example`, and `compose.yaml`:

| Variable | Default | Purpose |
| --- | --- | --- |
| `PASCALCALC_PORT` | `8088` | Host port for Docker Nginx gateway |
| `SERVER_PORT` | `8080` combined; `8081`/`8082`/`8083` in Compose | Go HTTP listener port |
| `MAX_REQUEST_BYTES` | `4096` | Maximum JSON request body size |
| `MAX_EXPRESSION_LENGTH` | `1024` | Backend expression length limit, in bytes |
| `BACKEND_API_URL` | `http://localhost:8080` | Vite development proxy target |
| `FRONTEND_PORT` | `5173` | Vite development server port |
| `VITE_API_BASE_URL` | `/api` | Browser API path prefix |
| `VITE_API_TIMEOUT_MS` | `10000` | Frontend API timeout setting |

**Important:** The Classic UI limit of 256 characters and backend limit of 1024 bytes are different. The backend's 1024-byte default is present in the uploaded Go configuration, `.env.example`, and Compose settings. Changing the backend default requires a coordinated implementation, test, configuration, and documentation update.

## 8. Testing and code coverage

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

## 9. Key design decisions and trade-offs

1. **Three services by capability:** Arithmetic, Math, and Expression are independently deployable and reusable across both UI modes. This adds infrastructure complexity compared with a monolith but demonstrates clear service boundaries.
2. **One Go module with separate entry points:** Shared interfaces and code reduce duplication, while standalone binaries permit independent container deployment. Shared-module changes may still affect multiple services.
3. **Nginx gateway:** Presents a single browser origin, centralizes routing, and avoids exposing internal Go service ports in Docker Compose.
4. **Centralized frontend API client:** Keeps transport details and endpoint configuration outside calculator UI components.
5. **Parser-based expression evaluation:** Tokenization and parsing support precedence, parentheses, unary operators, and functions without evaluating arbitrary source code.
6. **Layered validation:** UI input checks improve usability, while Go handlers and calculation services independently validate requests and numeric operations.
7. **Stateless JSON APIs:** No database is required for calculator operations. Go `float64` results have normal floating-point limitations; this is not an arbitrary-precision decimal engine.
8. **Testing and documentation first:** Unit tests, coverage, explicit API contracts, and phase-wise AI prompts support review and reproducibility.

## 10. Repository structure

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
├── .env.example
├── DOCKER_SETUP.md
└── README.md
```

## 11. Design documents and AI prompts — phase-wise

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
| **8. Deployment and Verification** | [Docker Setup](DOCKER_SETUP.md) and [Testing instructions](#8-testing-and-code-coverage) | Deploy the services, verify routes, run tests, and reproduce coverage |

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

## 12. Submission verification checklist

- [ ] `docker compose up --build -d` completes successfully.
- [ ] `http://localhost:8088` loads both calculator modes.
- [ ] Gateway `/health` and sample Arithmetic, Math, and Expression requests succeed.
- [ ] `go test ./... -count=1 -coverprofile=coverage.out` passes.
- [ ] `npm run test -- --run --coverage` passes.
- [ ] `npm run build` completes successfully.
- [ ] All relative documentation links resolve in the repository.
- [ ] Generated artifacts, local secrets, and unnecessary dependencies are excluded from Git.
- [ ] Final README matches the committed source, configuration, and documentation.

---

**Documentation note:** This README was prepared from the uploaded project archive and previously shared local coverage results. It is not a claim that Docker, Go tests, or frontend builds were executed in the README preparation environment.
