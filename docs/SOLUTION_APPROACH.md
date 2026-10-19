# PascalCalc — Solution Approach


## 1. Executive summary

PascalCalc — Classic Thinking. Modern Precision.

PascalCalc is a full-stack calculator application inspired by Blaise Pascal and the Pascaline mechanical calculator. It combines two complementary user experiences with a microservices-based Go backend.

The selected solution uses React, TypeScript, and Vite for the frontend; three independently deployable Go REST microservices for computation; Nginx for centralized API routing; and Docker for reproducible deployment.

Although a single modular backend would be simpler, the three-service approach has been selected to demonstrate service decomposition, independent deployment, API integration, modularity, and distributed-system engineering practices.

This document retrospectively consolidates the documented brainstorming decisions. It establishes the intended solution scope but does not finalize API schemas, internal package designs, infrastructure configuration, or detailed system architecture.

## 2. Problem statement and project objectives

### Problem statement

Develop a modern, responsive calculator that performs standard and extended mathematical operations while demonstrating professional full-stack engineering practices.

A basic calculator form would satisfy the minimum functional requirements but provide limited opportunities to demonstrate frontend state management, expression evaluation, API design, service decomposition, and deployment engineering.

PascalCalc addresses this by combining a traditional keypad-driven calculator with a structured calculation interface, supported by independently deployable backend capabilities.

### Project objectives

| ID     | Objective              | Expected evidence                                |
| ------ | ---------------------- | ------------------------------------------------ |
| OBJ-01 | Functional correctness | Seven supported operations                       |
| OBJ-02 | Frontend engineering   | Two responsive calculator modes                  |
| OBJ-03 | Backend engineering    | Three Go REST microservices                      |
| OBJ-04 | API design             | Versioned JSON contracts                         |
| OBJ-05 | Maintainability        | Clear responsibilities and reusable logic        |
| OBJ-06 | Reliability            | Validation, errors, and failure handling         |
| OBJ-07 | Testability            | Automated frontend and backend tests             |
| OBJ-08 | Deployment             | Reproducible Docker deployment                   |
| OBJ-09 | Documentation          | Solution, architecture, API, and testing records |

## 3. Proposed application concept

**PascalCalc — Classic Thinking. Modern Precision.**

Product vision: Deliver an engaging, accurate, and intuitive calculation experience through two interfaces powered by common mathematical capabilities.

**Classic Mode:** Numeric keypad, mathematical expressions, operator precedence, and parentheses.

**Precision Mode:** Explicit operands, selectable operations, and structured results.

The frontend modes are user-experience boundaries, not separate backend services. Both use shared API access and consistent mathematical behavior.

## 4. Selected features and scope

### 4.1 Functional requirements

| ID     | Requirement                   | Scope    |
| ------ | ----------------------------- | -------- |
| FR-001 | Addition                      | Included |
| FR-002 | Subtraction                   | Included |
| FR-003 | Multiplication                | Included |
| FR-004 | Division                      | Included |
| FR-005 | Exponentiation                | Included |
| FR-006 | Square root                   | Included |
| FR-007 | Percentage                    | Included |
| FR-008 | Classic keypad interface      | Included |
| FR-009 | Expression evaluation         | Included |
| FR-010 | Precision operation interface | Included |
| FR-011 | Result and error presentation | Included |
| FR-012 | Responsive layouts            | Included |

Classic Mode is expected to support constrained mathematical expressions, including parentheses and operator precedence. Precision Mode will use explicit numeric inputs, with the number of operands appropriate to each operation.

### 4.2 Non-functional requirements

| ID      | Requirement     | Proposed quality approach                     |
| ------- | --------------- | --------------------------------------------- |
| NFR-001 | Maintainability | Modular, cohesive components                  |
| NFR-002 | Testability     | Isolated units and API integration tests      |
| NFR-003 | Reliability     | Defined errors and graceful failure handling  |
| NFR-004 | Security        | Input constraints and safe expression parsing |
| NFR-005 | Performance     | Lightweight requests and bounded evaluation   |
| NFR-006 | Usability       | Responsive, accessible controls               |
| NFR-007 | Deployability   | Containerized independent services            |
| NFR-008 | Observability   | Health endpoints and structured logging       |
| NFR-009 | Compatibility   | Versioned REST API contracts                  |

Quantitative targets, such as response-time limits and minimum test coverage, remain to be agreed during architecture planning.

### 4.3 Application constraints

The following constraints capture PascalCalc input and error-handling boundaries. The frontend 256-character and backend 1024-byte limits were clarified and approved by the project owner during documentation review; they are not represented as requirements recorded in the original brainstorming prompts. Exact enforcement remains subject to implementation and test verification.

| ID | Constraint | Requirement | Evidence status |
| --- | --- | --- | --- |
| CON-001 | Classic Mode numeric input | Allow up to 15 digits per numeric operand | Reported; verify digit-counting and input behavior |
| CON-002 | Frontend expression length | Allow up to 256 characters per expression in the frontend | User-confirmed requirement; frontend enforcement pending verification |
| CON-003 | Backend expression length | Allow up to 1024 UTF-8 bytes per expression at the backend | User-confirmed requirement; backend byte counting and configuration to be verified |
| CON-004 | Division by zero | Reject division by zero with a meaningful error | Expected behavior; verify implementation and tests |
| CON-005 | Square root domain | Reject negative square-root inputs in real-number mode | Expected behavior; verify implementation and tests |
| CON-006 | Input validation | Reject invalid numeric inputs and unsupported expressions | Expected behavior; verify supported syntax and validation |
| CON-007 | Error handling | Present meaningful validation and calculation errors | Expected behavior; verify UI and API behavior |

Frontend character counting and backend UTF-8 byte counting are distinct limits. Detailed rules—such as handling leading zeros, decimals, pasted input, Unicode character counting, normalization, API status codes, and error payloads—belong in the frontend design, backend design, API reference, and testing documentation. Discrepancies discovered during verification must be recorded rather than silently changing the application.

### 4.4 Deferred and excluded features

Deferred: Calculation history, theme preferences, additional scientific functions, persistent storage, and advanced UI animations.

Excluded from the initial scope: Authentication, user accounts, databases without a demonstrated need, message queues, and Kubernetes.

## 5. Selected solution components

### Conceptual component view

```text
React + TypeScript + Vite
  ├── Classic Mode
  ├── Precision Mode
  └── Centralized API Client
                |
            HTTP / JSON
                |
              Nginx
        Central API routing
          /      |      \
 Arithmetic    Math    Expression
 Go REST      Go REST    Go REST
 +, -, *, /   Power, √, %  Parse / Evaluate
```

This is a conceptual view, not a verified deployment or API-routing diagram. Internal service-to-service communication remains unresolved at this phase.

### 5.1 Component responsibilities

| Component              | Primary responsibility                          |
| ---------------------- | ----------------------------------------------- |
| React frontend         | Presentation, interaction, and UI state         |
| Classic Mode           | Keypad and expression input                     |
| Precision Mode         | Structured operand and operation selection      |
| Centralized API client | HTTP communication and common response handling |
| Nginx                  | Request routing and frontend delivery           |
| Arithmetic Service     | Four basic arithmetic operations                |
| Math Service           | Power, square root, and percentage              |
| Expression Service     | Tokenization, parsing, and evaluation           |
| Docker deployment      | Independent service packaging and orchestration |

Each Go service is intended to have its own runtime and deployment artifact. Verification against the implementation belongs to later phases.

The Expression Service's relationship with the Arithmetic and Math Services requires special attention. It may use REST calls to those services or a shared, versioned calculation library. This decision is intentionally left open.

## 6. Architecture alternatives considered

| Alternative                               | Advantages                                                            | Disadvantages                                | Decision     |
| ----------------------------------------- | --------------------------------------------------------------------- | -------------------------------------------- | ------------ |
| A — Single modular Calculator Service     | Simplest deployment, shared logic, easy testing                       | One deployment boundary                      | Not selected |
| B — Classic and Precision Services        | Independent deployment by UI mode                                     | Duplicated mathematics, UI-driven boundaries | Rejected     |
| C — Arithmetic, Math, Expression Services | Capability ownership, independent deployment, distributed integration | Higher networking and operational complexity | Selected     |

### Rationale for Option C

The three-service approach aligns backend boundaries with mathematical capabilities rather than frontend presentation modes.

It provides opportunities to demonstrate independent Go service development, REST contracts, containerized deployment, routing, integration testing, and service-level error handling.

However, it is important to acknowledge that this is an project-driven architectural choice, not a demonstrated production scaling requirement. The mathematical workload alone does not justify three distributed services.

The selected architecture should therefore remain deliberately lightweight, with no unnecessary infrastructure.

## 7. Key architectural assumptions and trade-offs

| ID    | Assumption or trade-off                                          | Implication                                        |
| ----- | ---------------------------------------------------------------- | -------------------------------------------------- |
| AS-01 | Services are independently deployable                            | Separate Go executables and containers             |
| AS-02 | REST/JSON is the integration mechanism                           | Clear contracts but serialization/network overhead |
| AS-03 | Nginx is the central routing layer                               | Simplifies frontend service discovery              |
| AS-04 | Calculation requests are stateless                               | No initial database dependency                     |
| AS-05 | Mathematical logic should not be duplicated                      | Requires a deliberate reuse strategy               |
| AS-06 | Expression evaluation is bounded                                 | Grammar and resource limits required               |
| AS-07 | Docker is the deployment baseline                                | Reproducible local development environment          |
| AS-08 | Distributed complexity is acceptable for project goals and engineering practice | More integration and operational tests             |

Key trade-off: PascalCalc gains independently deployable capability boundaries at the expense of greater operational complexity and possible network dependencies.

The architecture review must ensure that the service boundaries remain meaningful and that the Expression Service does not become unnecessarily coupled to the other services.

## 8. Risks and mitigation strategies

| Risk                           | Impact                           | Proposed mitigation                                       |
| ------------------------------ | -------------------------------- | --------------------------------------------------------- |
| Distributed call failures      | Failed expressions               | Timeouts, structured errors, appropriate failure handling |
| Mathematical logic duplication | Inconsistent results             | Shared domain contracts and a deliberate reuse approach   |
| Expression parser defects      | Incorrect calculations           | Defined grammar and comprehensive parser tests            |
| Floating-point behavior        | Unexpected decimal results       | Document precision and formatting rules                   |
| API contract changes           | Frontend integration failures    | Versioning and contract tests                             |
| Excessive service complexity   | Delivery delays                  | Limit infrastructure and avoid speculative features       |
| Incorrect Nginx routing        | API unavailability               | Routing tests and deployment smoke tests                  |
| Inconsistent validation        | Different behavior between modes | Shared validation rules and cross-mode test cases         |
| Unbounded expressions          | Resource exhaustion              | Input length, depth, and computational limits             |

Retry behavior must be considered carefully. Calculations are expected to be stateless, but indiscriminate retries can amplify load during failures. Any retry policy should be bounded and justified

## 9. Questions for formal high-level design

- What are the exact REST endpoint paths, JSON schemas, and HTTP error contracts?
- Will the Expression Service reuse calculation logic through a shared library or invoke other services over HTTP?
- What expression grammar, length, depth, and numeric constraints apply?
- How will timeouts, failures, logging, and health checks be implemented?
- What are the Nginx routing and Docker Compose networking conventions?
- Which testing and coverage targets will be adopted and measured?

## 10. Criteria for proceeding to architecture design

- The project scope, calculator modes, and seven operations are agreed.
- The three-service selection and its operational trade-offs are accepted.
- Deferred and excluded features are recorded.
- Unresolved design questions are assigned to the architecture phase rather than assumed resolved.
- This document is reviewed and explicitly approved as the Phase 2 input.

**Approval status:** Pending. No approval or implementation verification is implied by this document.
