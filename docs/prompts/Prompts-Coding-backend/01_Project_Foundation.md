# PascalCalc — Backend Coding Prompts: 01 Project Foundation

## 1. Purpose

Establish the Go module, shared API identifiers, runtime configuration, and unit tests for configuration logic needed by the Arithmetic, Math, and Expression services. Complete this foundation before developing domain logic or HTTP handlers.

## 2. Inputs and dependencies

- `Docs/SOLUTION_APPROACH.md` — project requirements: 256 frontend characters and 1024 backend expression bytes.
- `Docs/architecture.md` — three independent REST services, stateless requests, and gateway routing.
- `Docs/backend-hld.md` — Go module structure, common contracts, configuration, and validation strategy.
- Go 1.19, as declared by the uploaded `backend/go.mod`.
- No domain service or HTTP handler implementation is a prerequisite.

**Dependency sequence within this component:** Go module and package layout → API constants → configuration loading and validation → service-specific configuration → environment-file support → unit tests → review.

## 3. Main Coding Prompt — P04-01-01



> Act as a senior Go backend engineer building the foundation of PascalCalc. Follow the approved Solution Approach, System Architecture, and Backend HLD. Work only within `backend/`. Do not implement arithmetic, math, expression parsing, HTTP handlers, routers, service startup, or Docker configuration in this step.
>
> 1. Initialize the Go module as `calculator-app/backend` with Go 1.19 compatibility. Keep dependencies minimal; use `github.com/joho/godotenv` for optional `.env` loading if required by the design.
> 2. Establish `internal/api` with named constants for the versioned Arithmetic, Math, and Expression routes, `/health`, and the existing legacy `/api/v1/calculate` compatibility route. Define typed operation identifiers for add, subtract, multiply, divide, power, square root, and percentage. Keep route names and operation identifiers centralized.
> 3. Establish `internal/config` with a typed `Config` containing server host/port, maximum request bytes, maximum expression length, and HTTP read/write/idle timeouts. Load values from environment variables with explicit defaults and return descriptive errors for malformed or invalid settings. Validate port range, positive size limits, nonempty host, and positive timeouts.
> 4. Provide a `LoadForService(defaultPort int)` function so each service can use its own default port while respecting an explicit `SERVER_PORT` override. Avoid introducing any service-to-service network dependency.
> 5. Provide optional `.env` loading with existing operating-system environment variables taking precedence. Missing optional `.env` files should not prevent containerized execution; malformed or unreadable files should return an error.
> 6. Set the backend `MAX_EXPRESSION_LENGTH` default to **1024 bytes**, consistent with the approved backend requirement. This is a byte limit, not a character limit; Go `len(string)` measures UTF-8 bytes. The separate **256-character frontend limit** is outside this component. Do not change existing application code/configuration in the documentation-validation workflow without approval.
> 7. Return the proposed file tree, source changes, design decisions, and build commands. Do not modify existing source without explicit approval in this documentation-validation workflow.

**Expected source mapping (observed in uploaded ZIP):**

| Responsibility | Source file |
|---|---|
| Go module and dependency declaration | `backend/go.mod`, `backend/go.sum` |
| Versioned routes and health path | `backend/internal/api/routes.go` |
| Operation identifiers | `backend/internal/api/operations.go` |
| General config loading and validation | `backend/internal/config/config.go` |
| Per-service default ports | `backend/internal/config/service.go` |
| Optional environment file | `backend/internal/config/dotenv.go` |

## 4. Unit Testing Prompt — P04-01-02



> Act as a Go unit-test engineer. Use the implementation and design decisions produced by P04-01-01. Create isolated, table-driven unit tests for `internal/config` and, where appropriate, pure API identifier definitions. Test package functions and configuration behavior without starting HTTP servers, containers, or external services.
>
> Cover: defaults; valid environment overrides; malformed integer, integer64, and duration values; invalid port boundaries (0, 65536); empty host where applicable; nonpositive request/expression limits; nonpositive timeouts; service-specific default ports; explicit `SERVER_PORT` precedence; missing optional `.env`; and environment precedence. Test the backend default of **1024 bytes**, including environment overrides and invalid limits; do not assert a 256-byte backend default. Restore environment variables between test cases to prevent leakage.
>
> Run focused unit tests using `cd backend && go test ./internal/config ./internal/api` and measure component coverage with `cd backend && go test -cover ./internal/config ./internal/api` if the Go toolchain is available. Report exact commands, observed results, failing cases, and coverage only when measured. Do not invent test execution or coverage figures. Exclude API integration, gateway, Docker, deployment, and end-to-end testing from this prompt. Do not edit existing implementation to force tests to pass without approval.


- `backend/internal/config/config_test.go`
- `backend/internal/config/service_test.go`
- No dedicated `internal/api` test file was observed in the inspected ZIP; do not claim one exists.

## 5. Code Review Prompt — P04-01-03



> Act as a principal Go engineer. Review the implementation from P04-01-01 and the unit-test code and results from P04-01-02 against the approved prior phase documents. Check module/package cohesion, import hygiene, Go 1.19 compatibility, route/operation naming consistency, configuration precedence, error messages, range validation, environment isolation in tests, dependency minimization, and absence of domain logic or HTTP server responsibilities in the foundation.
>
> Verify that the **1024-byte backend** configuration default is consistent with the approved requirement and that the separate **256-character frontend** limit is not incorrectly applied to backend configuration. Classify findings as defects, documentation discrepancies, or unverified behavior. For each finding provide severity, evidence (file/function), proposed remediation, and whether user approval is required. Do not apply corrections automatically.
>
> Deliver a review checklist, prompt-to-file mapping, unit-test evidence summary, unresolved decisions, and a **Pass / Partial / Fail / Not verified** outcome. A static source inspection is not equivalent to successful test execution.

## 6. Traceability and validation status

| Prompt ID | Output / evidence | Current validation |
|---|---|---|
| P04-01-01 | `go.mod`, `internal/api/*.go`, `internal/config/{config,service,dotenv}.go` | Source files located; prompt is reconstructed, not executed |
| P04-01-02 | `internal/config/{config,service}_test.go` | Unit-test files located; execution not verified |
| P04-01-03 | Review checklist and discrepancy register | Prompt drafted; review not executed |

**Clarification FND-01 (resolved):** The existing `MaxExpressionLength: 1024` numeric default aligns with the owner-approved **1024-byte backend** limit. The frontend limit is **256 characters**. The earlier P04-01-01 isolated re-execution generated a 256 default under the superseded interpretation; that generated result is **not conformant** with the clarified backend requirement and must be regenerated or corrected only in the isolated validation workspace. The historical validation report must not be read as evidence that the 256 default is correct.

**Completion gate:** Review this prompt document, validate the foundation implementation and tests, resolve or formally defer open discrepancies, and approve Component 01 before proceeding to Component 02 — Arithmetic.

## 7. Related documents

- [Solution Approach](../../../Docs/SOLUTION_APPROACH.md)
- [System Architecture](../../../Docs/architecture.md)
- [Backend High-Level Design](../../../Docs/backend-hld.md)
