# Prompt 1 — Initial three-microservice system design
Act as a senior full-stack solution architect. Design PascalCalc, a full-stack calculator application with a React + TypeScript (Vite) frontend and three independently runnable and deployable Go REST microservices:
1. Arithmetic Service — addition, subtraction, multiplication, division.
2. Math Service — exponentiation, square root, percentage.
3. Expression Service — parsing and evaluating calculator expressions, including operator precedence.
The frontend offers Classic Mode (keypad and expression entry) and Precision Mode (operand inputs and operation buttons). Design a single browser-facing API entry point using Nginx as a reverse proxy; retain the public paths /api/v1/arithmetic/*, /api/v1/math/*, and /api/v1/expressions/evaluate. Do not duplicate mathematical evaluation in React.
Provide: architecture and deployment-boundary diagrams; service responsibilities; proposed Go project structure; API request/response contracts; validation and error handling; independent configuration and health checks; test strategy; and initial design decisions. Prefer idiomatic Go, simple deployment, and minimal dependencies. Clearly label planned rather than implemented features. Do not generate implementation code.
Expected review artifact: Initial three-service HLD with assumptions and open questions.
# Prompt 2 — Review the initial architecture
Act as a principal engineer reviewing the Prompt 1 three-microservice HLD against the assessment requirements: clean React UI, Go REST APIs, validation, JSON results, maintainability, tests and coverage, documentation, and optional Docker.
Challenge the design on:
- Whether three deployable processes are genuinely independent (startup, routing, ports, health checks, shutdown, Docker images).
- Whether one Go module and shared internal packages undermine deployment independence, and whether separate Go modules are justified.
- Nginx gateway routing, stable browser-facing API URLs, local development without Docker, and avoiding browser CORS problems.
- Request validation, finite-number handling, division by zero, negative square roots, overflow, expression-length limits, and consistent error contracts.
- Expression evaluation state isolation, timeouts, graceful shutdown, and failure behavior when one service is unavailable.
- Unit, handler, gateway, integration, and end-to-end test coverage.
- Operational trade-offs of three services for a small calculator domain.
Identify issues with severity, propose concrete changes, compare alternatives, and record accepted/rejected recommendations. Do not claim that the services already exist. Do not generate code.
Expected review artifact: Review findings and change decisions to feed Prompt 3.
# Prompt 3 — Revised HLD after architecture review
Act as a solution architect. Using the initial HLD from Prompt 1 and the approved review findings from Prompt 2, produce the revised PascalCalc architecture for three independently deployable Go microservices, without changing the frontend's public API contracts.
Specify:
1. Service ownership and isolation: Arithmetic, Math, Expression.
2. Separate cmd/arithmetic-api, cmd/math-api, and cmd/expression-api entry points, routers, health endpoints, runtime configuration, and shutdown lifecycle.
3. A single Go module with reusable internal code as the initial implementation choice; explicitly discuss compile-time sharing versus runtime coupling.
4. Nginx gateway mapping for /api/v1/arithmetic/, /api/v1/math/, and /api/v1/expressions/ while preserving the original paths.
5. JSON contracts, request limits, validation rules, error responses, service unavailability, and edge cases.
6. React Classic/Precision request flows, including the expression evaluator and square-root endpoint.
7. Mermaid architecture, request-sequence, and deployment diagrams; directory layout; local-run and Docker plans.
8. Verification criteria, tests, coverage expectations, risks, and trade-offs.
Produce docs/architecture.md and docs/backend-hld.md as design documents. Clearly distinguish target architecture from verified implementation status. Do not generate implementation code.
Expected review artifact: Revised HLD ready for final design validation.
# Prompt 4 — Final assessment readiness and design sign-off
Act as an engineering design review board. Review the revised three-service HLD from Prompt 3 against the original assessment and the UI blueprint. Verify that:
- All three backend services can be built, started, stopped, tested, and deployed independently.
- API paths and JSON payloads remain compatible with both React calculator modes.
- Nginx routes each endpoint to the correct service and returns appropriate errors for unavailable upstreams.
- Backend handlers enforce validation and error handling independently of frontend validation.
- Unit tests, handler tests, integration tests, and coverage reports are planned and measurable.
- Docker Compose can run the frontend, gateway, and all three Go services with health checks.
- Documentation includes setup, API examples, design rationale, and explicit assessment traceability.
Record Architecture Decision Records for (1) three service boundaries, (2) Nginx gateway, and (3) shared Go module versus separate modules. Provide an acceptance checklist, unresolved risks, and implementation milestones. Do not mark unimplemented items as complete. Do not generate implementation code.# 