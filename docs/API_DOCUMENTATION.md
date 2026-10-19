# PascalCalc — Backend API Documentation

**Phase:** 5 — API Documentation  
**Status:** Source-derived API reference (not an end-to-end deployment certification)  
**Source baseline:** `calculator-app(2).zip`, Go backend handlers and routers  
**API version:** v1

## Overview

PascalCalc exposes JSON REST endpoints through three independently deployable Go services: Arithmetic, Math, and Expression. A combined legacy server entry point is also present. In the intended deployment, Nginx forwards the versioned paths to the appropriate service; direct service requests use the same endpoint paths.

| Service | Default standalone port | Operations |
|---|---:|---|
| Arithmetic | 8081 | Add, subtract, multiply, divide |
| Math | 8082 | Power, square root, percentage |
| Expression | 8083 | Evaluate mathematical expressions |
| Combined server | 8080 | All eight calculation routes and health |

**Note:** Ports are defaults. `SERVER_PORT` can override them. The API paths below are defined by the Go application; Nginx deployment routing is outside the scope of this source-derived contract.

## Request and response conventions

- Calculation routes accept **POST** with a single JSON object. `Content-Type: application/json` is expected by clients; handlers parse JSON without explicitly enforcing the request content-type header.
- Successful calculation responses return HTTP **200** and JSON. Invalid input or calculation errors return HTTP **400** and `{"error":"message"}`.
- Unsupported methods return HTTP **405** with an `Allow: POST` header on calculation routes.
- JSON decoding rejects unknown fields and additional JSON values. Missing required numeric fields are rejected.
- The configured request body limit defaults to **4096 bytes** (`MAX_REQUEST_BYTES`). Exceeding the limit is handled as a JSON decode error with HTTP 400, not a dedicated 413 response.
- Results are JSON numbers based on Go `float64`; clients should not assume arbitrary precision.

## Arithmetic Service

All four routes accept `{"a": number, "b": number}`. Both operands are required and must be finite. A successful response is `{"result": number}`. The handler rejects non-finite arithmetic results.

| Method | Endpoint | Operation | Example input | Example output |
|---|---|---|---|---|
| POST | `/api/v1/arithmetic/add` | a + b | `{"a":12,"b":3}` | `{"result":15}` |
| POST | `/api/v1/arithmetic/subtract` | a − b | `{"a":12,"b":3}` | `{"result":9}` |
| POST | `/api/v1/arithmetic/multiply` | a × b | `{"a":12,"b":3}` | `{"result":36}` |
| POST | `/api/v1/arithmetic/divide` | a ÷ b | `{"a":12,"b":3}` | `{"result":4}` |

**Representative HTTP 400 errors:** `invalid JSON request`, `unexpected JSON content`, `both operands are required`, `operands must be finite numbers`, `cannot divide by zero`, `invalid calculation result`.

### Arithmetic example

```bash
curl -X POST http://localhost:8081/api/v1/arithmetic/divide \
  -H 'Content-Type: application/json' \
  -d '{"a":12,"b":3}'
```

Response (200):

```json
{"result":4}
```

## Math Service

The Math Service uses numeric `a` and `b` fields. `a` is required for all operations; `b` is required for power and percentage, but not square root. The square-root handler accepts a supplied `b` field but does not use it. Success returns `{"result": number}`.

| Method | Endpoint | Semantics | Example input | Example output |
|---|---|---|---|---|
| POST | `/api/v1/math/power` | a raised to b | `{"a":2,"b":3}` | `{"result":8}` |
| POST | `/api/v1/math/sqrt` | square root of a | `{"a":81}` | `{"result":9}` |
| POST | `/api/v1/math/percentage` | a percent of b | `{"a":20,"b":150}` | `{"result":30}` |

**Representative HTTP 400 errors:** `invalid JSON request`, `unexpected JSON content`, `operand a is required`, `operand b is required`, `inputs must be finite numbers`, `input must be a finite number`, `cannot calculate square root of a negative number`, `invalid exponentiation result`, `invalid percentage result`.

The Math handler delegates finite-number validation to the underlying service. The service rejects invalid mathematical results such as an infinite power result.

### Math example

```bash
curl -X POST http://localhost:8082/api/v1/math/percentage \
  -H 'Content-Type: application/json' \
  -d '{"a":20,"b":150}'
```

Response (200):

```json
{"result":30}
```

## Expression Service

| Method | Endpoint | Request | Success response |
|---|---|---|---|
| POST | `/api/v1/expressions/evaluate` | `{"expression":"2+3*4"}` | `{"expression":"2+3*4","result":14}` |

The handler trims leading/trailing whitespace before evaluation and returns the **trimmed** expression in its response. An empty expression is rejected. The default backend expression limit is **1024 UTF-8 bytes** (`MAX_EXPRESSION_LENGTH`); the evaluator checks Go string byte length. The frontend's separate 256-character input limit does not replace the backend limit.

The expression parser supports arithmetic, parentheses, unary plus/minus, exponentiation, and square-root syntax as defined by its tokenizer and evaluator. For exact accepted symbols and precedence, refer to `internal/expression/tokenizer.go` and `evaluator.go`.

**Representative HTTP 400 errors:** `invalid JSON request`, `unexpected JSON content`, `expression is required`, `invalid expression encoding`, `expression too long`, `invalid calculation result`, and tokenizer/parser-specific messages.

### Expression example

```bash
curl -X POST http://localhost:8083/api/v1/expressions/evaluate \
  -H 'Content-Type: application/json' \
  -d '{"expression":"2+3*4"}'
```

Response (200):

```json
{"expression":"2+3*4","result":14}
```

## Sample JSON request and response payloads

The following payloads show the **POST request body** and the corresponding **HTTP 200 response body** for every implemented calculation endpoint. They are illustrative examples derived from the source contract, not results from live HTTP tests.

### Arithmetic endpoints

**POST `/api/v1/arithmetic/add`**

Request:
```json
{"a": 12, "b": 3}
```
Response (200):
```json
{"result": 15}
```

**POST `/api/v1/arithmetic/subtract`**

Request:
```json
{"a": 12, "b": 3}
```
Response (200):
```json
{"result": 9}
```

**POST `/api/v1/arithmetic/multiply`**

Request:
```json
{"a": 12, "b": 3}
```
Response (200):
```json
{"result": 36}
```

**POST `/api/v1/arithmetic/divide`**

Request:
```json
{"a": 12, "b": 3}
```
Response (200):
```json
{"result": 4}
```

### Math endpoints

**POST `/api/v1/math/power`**

Request:
```json
{"a": 2, "b": 3}
```
Response (200):
```json
{"result": 8}
```

**POST `/api/v1/math/sqrt`**

Request:
```json
{"a": 81}
```
Response (200):
```json
{"result": 9}
```

**POST `/api/v1/math/percentage`** — calculates `a` percent of `b`.

Request:
```json
{"a": 20, "b": 150}
```
Response (200):
```json
{"result": 30}
```

### Expression endpoint

**POST `/api/v1/expressions/evaluate`**

Request:
```json
{"expression": "2+3*4"}
```
Response (200):
```json
{"expression": "2+3*4", "result": 14}
```

### Health endpoint

**GET `/health`** — no request body.

Response (200):
```json
{"status": "ok"}
```

### Sample error payloads

**Division by zero — POST `/api/v1/arithmetic/divide`**

Request:
```json
{"a": 12, "b": 0}
```
Response (400):
```json
{"error": "cannot divide by zero"}
```

**Missing operand — POST `/api/v1/arithmetic/add`**

Request:
```json
{"a": 12}
```
Response (400):
```json
{"error": "both operands are required"}
```

**Negative square root — POST `/api/v1/math/sqrt`**

Request:
```json
{"a": -9}
```
Response (400):
```json
{"error": "cannot calculate square root of a negative number"}
```

**Missing expression — POST `/api/v1/expressions/evaluate`**

Request:
```json
{"expression": ""}
```
Response (400):
```json
{"error": "expression is required"}
```

## Health endpoint

| Method | Endpoint | Success | Incorrect method |
|---|---|---|---|
| GET | `/health` | 200, `{"status":"ok"}` | 405, plain-text `method not allowed` |

The health endpoint is registered on all three standalone service routers and the combined router. Unlike calculation endpoints, its 405 response is generated using `http.Error` and is **not JSON**.

## Errors and status codes

| Status | Meaning | Body format |
|---|---|---|
| 200 | Successful calculation or health check | Endpoint-specific JSON |
| 400 | Invalid JSON, validation failure, or calculation error | `{"error":"..."}` |
| 405 | Unsupported HTTP method | JSON for calculation routes; plain text for health |
| 404 | Unregistered route | Default Go `net/http` response; not the application JSON error schema |

There is no evidence in the inspected handlers of application-defined authentication, authorization, rate limiting, or a dedicated 429/413 contract. Such behavior could be added by a deployment proxy, but is not documented here as implemented.

## Service routing and compatibility

| Application entry point | Registered routes |
|---|---|
| `cmd/arithmetic-api/main.go` | Four Arithmetic routes + `/health` |
| `cmd/math-api/main.go` | Three Math routes + `/health` |
| `cmd/expression-api/main.go` | Expression evaluation + `/health` |
| `cmd/api/main.go` | All eight calculation routes + `/health` |

**Legacy route warning:** `internal/api/routes.go` defines `/api/v1/calculate` as `LegacyCalculate`, but no corresponding handler registration exists in the inspected Go source. It should **not** be advertised as available until implemented or confirmed in a different deployed version.

## Configuration affecting API behavior

| Environment variable | Default | Effect |
|---|---|---|
| `SERVER_HOST` | `0.0.0.0` | Listen address |
| `SERVER_PORT` | `8080` in combined mode; standalone defaults 8081/8082/8083 | Listen port |
| `MAX_REQUEST_BYTES` | `4096` | Maximum request body size |
| `MAX_EXPRESSION_LENGTH` | `1024` | Maximum expression byte length |
| `HTTP_READ_TIMEOUT` | `5s` | HTTP read timeout |
| `HTTP_WRITE_TIMEOUT` | `10s` | HTTP write timeout |
| `HTTP_IDLE_TIMEOUT` | `60s` | HTTP idle timeout |

## Source references

- `backend/internal/api/routes.go` — route constants
- `backend/internal/arithmetic/handler.go`, `service.go` — arithmetic contract and errors
- `backend/internal/mathops/handler.go`, `service.go` — math contract and errors
- `backend/internal/expression/handler.go`, `tokenizer.go`, `evaluator.go` — expression contract
- `backend/internal/server/services.go`, `router.go` — service and combined routing
- `backend/internal/config/config.go`, `service.go` — defaults and overrides
- `backend/cmd/{arithmetic-api,math-api,expression-api,api}/main.go` — executable entry points

## Documentation limitations

This reference is derived from inspected source code, not a running deployed environment. Examples illustrate the implementation contract; they are not claimed as observed HTTP responses from an end-to-end integration test. Any API changes should be reflected here after implementation and verification.
