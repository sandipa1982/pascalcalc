# Prompt 1 — Requirements Analysis and Initial Brainstorming

Act as a Senior Full-Stack Engineer and Solution Architect.

I have a technical assessment to build a full-stack calculator application with a React frontend and a backend microservice. React with TypeScript and Go are the preferred technologies.

The application must support addition, subtraction, multiplication, and division. Optional operations include exponentiation, square root, and percentage.

The assessment emphasizes clean design, maintainable code, testability, input validation, error handling, responsive UI, and documentation.

Before designing the architecture, help me brainstorm ideas for building a calculator application that meets these requirements and demonstrates strong engineering practices.

Explore:

Core functional and non-functional requirements.

Potential application features and user experiences.

Different ways to organize the frontend and backend.

Opportunities to demonstrate modularity and microservices.

Testing, reliability, and deployment considerations.

Features that would add value without unnecessary complexity.

Recommend a practical project scope and identify decisions requiring further discussion.

Do not generate implementation code or finalize the architecture yet.

# Prompt 2 — Product Concept and Feature Brainstorming

Based on the initial requirements analysis, I want to develop a calculator application that is more engaging than a basic arithmetic form while remaining within the technical assessment scope.

Explore the idea of combining a traditional calculator experience with a more structured calculation interface.

Consider:

A Classic Calculator with a numeric keypad and expression-based calculations.

A Precision Calculator with explicit numeric inputs and selectable mathematical operations.

Shared calculation capabilities exposed through REST APIs.

Clear result presentation, validation messages, and responsive layouts.

Support for all seven mathematical operations.

Evaluate the usability, implementation complexity, maintainability, and testing implications of these ideas.

Also explore a product identity inspired by the history of mechanical calculators, particularly Blaise Pascal and the Pascaline.

Propose a suitable application name, tagline, and high-level feature scope.

Explain which ideas should be accepted, deferred, or rejected.

Do not create wireframes or implementation code yet.
# Prompt 3 — Backend Architecture Alternatives and Trade-offs

Act as a Principal Software Architect.

For the PascalCalc application concept, evaluate different ways to organize the Go backend as independently deployable REST microservices.

Consider these alternatives:

Option A — Single Calculator Microservice
One backend service with internal arithmetic, mathematical, and expression-evaluation modules.

Option B — Two Calculator-Mode Microservices
A Classic Calculator service and a Precision Calculator service.

Option C — Three Capability-Based Microservices

Arithmetic Service: addition, subtraction, multiplication, division.

Math Service: power, square root, percentage.

Expression Service: expression tokenization and evaluation.

Compare the alternatives using:

Separation of concerns and service ownership.

Independent deployment and scalability.

Code reuse and duplication.

Coupling between frontend and backend.

REST API design and compatibility.

Unit and integration testing.

Docker deployment and operational complexity.

Suitability for the technical assessment.

Recommend the most appropriate architecture, explain the trade-offs, and identify risks that should be addressed during the formal architecture design review.

Do not generate code or finalize the HLD yet.

# Prompt 4 — PascalCalc Solution Selection and Design Readiness

Act as a Senior Engineering Technical Lead.

Based on the previous brainstorming discussions, consolidate the proposed PascalCalc solution.

The selected approach is:

Frontend

React with TypeScript and Vite.

Classic and Precision calculator modes.

A centralized API client.

Backend

Three independently deployable Go microservices: Arithmetic, Math, and Expression.

Versioned REST endpoints with JSON requests and responses.

Centralized API routing through Nginx.

Reusable calculation logic and clear module responsibilities.

Quality and deployment

Input validation and error handling.

Frontend and backend unit tests.

Docker-based deployment.

Architecture documentation and design decisions.

Prepare a solution proposal containing:

Problem statement and assessment objectives.

Proposed PascalCalc application concept.

Selected features and scope.

Alternatives considered and reasons for selecting the three-service architecture.

High-level component responsibilities.

Key architectural assumptions and trade-offs.

Identified risks and mitigation strategies.

Questions to resolve during formal HLD.

Acceptance criteria for proceeding to architecture design.

Clearly distinguish brainstorming decisions from finalized technical designs.

Do not generate implementation code. The next phase will be the formal system architecture design.