# PascalCalc — Frontend Coding Prompts: 01 Project Foundation



## Purpose and scope

Build the Project Foundation component without changing the original PascalCalc application. This component is responsible only for the files listed below; keep concerns owned by other components separate.

**Existing-source mapping:**
- `frontend/package.json`
- `frontend/vite.config.ts`
- `frontend/tsconfig.json`
- `frontend/tsconfig.app.json`
- `frontend/tsconfig.node.json`
- `frontend/index.html`
- `frontend/src/main.tsx`
- `frontend/src/index.css`

## Prerequisites and dependencies

-  `FRONTEND_DESIGN.md`, approved `FRONTEND_BLUEPRINT`, and`API_DOCUMENTATION.md`.
- React + TypeScript + Vite toolchain matching the current `frontend/package.json`.
- Execute in an isolated working copy; do not modify the original project without approval.
- Preserve the existing API endpoint contracts and architecture boundaries.

## Main Coding Prompt — P07-01-01


> Act as a senior React/TypeScript frontend engineer building PascalCalc , Component 01 (Project Foundation). Use `FRONTEND_BLUEPRINT`, `FRONTEND_DESIGN.md`, API documentation, and the approved prior-phase architecture. Reconstruct the following existing implementation responsibilities in an isolated workspace.
>
> 1. Work only on the component's mapped files: `frontend/package.json`, `frontend/vite.config.ts`, `frontend/tsconfig.json`, `frontend/tsconfig.app.json`, `frontend/tsconfig.node.json`, `frontend/index.html`, `frontend/src/main.tsx`, `frontend/src/index.css`.
> 2. Set up React, TypeScript and Vite with Vitest and React Testing Library; preserve package scripts, strict TypeScript settings, application mounting, global styles and the existing development proxy. Keep all calculations on the Go backend.
> 3. Preserve the UI design's accessible controls, clear loading and error feedback, and responsive layout where applicable. Distinguish blueprint targets from observed source behavior; flag ambiguities rather than silently changing behavior.
> 4. Keep all mathematical calculations and expression precedence on the Go backend. Do not create new endpoints, duplicate Go calculation logic, or change unrelated components.
> 5. Return complete generated source, design assumptions, and the exact formatting/typecheck/build commands. Only claim command success if commands were executed.

## Unit Testing Prompt — P07-01-02

**Type:** Reconstructed unit-testing prompt

> Act as a frontend unit-test engineer. Use the code produced by P07-01-01. Write focused Vitest and React Testing Library tests where appropriate.
>
> 1. Validate app bootstrap, TypeScript configuration and test runner using smoke/unit checks; do not add browser end-to-end or integration tests.
> 2. Test accessible labels, visible state and errors where applicable. Avoid fragile implementation-detail assertions.
> 3. Mock external modules, network access and timers when needed; these are **unit tests**, not backend integration tests or browser E2E tests.
> 4. Run `npm test -- --run` or the actual test script defined in `package.json`, plus TypeScript checks if available; report actual results only.
> 5. Do not alter application source merely to make tests pass without approval.

## Code Review Prompt — P07-01-03

**Type:** Reconstructed code-review prompt

> Act as a principal frontend engineer. Review P07-01-01 implementation and P07-01-02 unit tests against the approved Phase 6 design and the existing PascalCalc source.
>
> 1. Check functionality, component boundaries, API-contract alignment, accessibility, state transitions, validation and error handling.
> 2. Check test quality and coverage of normal, boundary and failure cases. Confirm there are no unintended live-network dependencies in unit tests.
> 3. Compare generated source with existing mapped files for **behavioral equivalence**, not byte-for-byte identity; identify any intentional differences explicitly.
> 4. Run formatting/lint/typecheck/test/build commands only if feasible, recording actual output. Mark blocked checks **Not verified**.
> 5. Write findings separately in the validation workspace with Pass / Partial / Fail / Not verified and recommended actions. Do not change original application code without approval.

## Related documents

- [Frontend Design](../../../Docs/FRONTEND_DESIGN.md)
- [UI Design Blueprint](../../../Docs/FRONTEND_BLUEPRINT.md)
- [API Documentation](../../../Docs/API_DOCUMENTATION.md)
- [System Architecture](../../../Docs/architecture.md)
