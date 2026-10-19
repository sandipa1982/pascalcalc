# PascalCalc — Frontend Coding Prompts: 04 Classic Calculator



## Purpose and scope

Build the Classic Calculator component without changing the original PascalCalc application. This component is responsible only for the files listed below; keep concerns owned by other components separate.

**Existing-source mapping:**
- `frontend/src/components/ClassicCalculator.tsx`
- `frontend/src/components/ClassicCalculator.css`
- `frontend/src/components/ClassicCalculator.test.tsx`

## Prerequisites and dependencies

- `FRONTEND_DESIGN.md`, approved `FRONTEND_BLUEPRINT`, and  `API_DOCUMENTATION.md`. Complete Component 03 and its relevant dependencies first.
- React + TypeScript + Vite toolchain matching the current `frontend/package.json`.
- Execute in an isolated working copy; do not modify the original project without approval.
- Preserve the existing API endpoint contracts and architecture boundaries.

## Main Coding Prompt — P07-04-01

**Type:** Reconstructed main coding prompt

> Act as a senior React/TypeScript frontend engineer building PascalCalc , Component 04 (Classic Calculator). Use `FRONTEND_BLUEPRINT`, `FRONTEND_DESIGN.md`, API documentation, and the approved prior-phase architecture. Reconstruct the following existing implementation responsibilities in an isolated workspace.
>
> 1. Work only on the component's mapped files: `frontend/src/components/ClassicCalculator.tsx`, `frontend/src/components/ClassicCalculator.css`, `frontend/src/components/ClassicCalculator.test.tsx`.
> 2. Implement the Classic keypad, display, expression editing, operators, parentheses, clear, sign, percentage, power, square root, completed expression and loading/error states. Preserve MAX_CLASSIC_EXPRESSION_LENGTH = 256 JavaScript characters. Use evaluateExpression for expression evaluation and the existing source-defined behavior for special keys; verify square-root wiring rather than assuming the blueprint is already implemented.
> 3. Preserve the UI design's accessible controls, clear loading and error feedback, and responsive layout where applicable. Distinguish blueprint targets from observed source behavior; flag ambiguities rather than silently changing behavior.
> 4. Keep all mathematical calculations and expression precedence on the Go backend. Do not create new endpoints, duplicate Go calculation logic, or change unrelated components.
> 5. Return complete generated source, design assumptions, and the exact formatting/typecheck/build commands. Only claim command success if commands were executed.

## Unit Testing Prompt — P07-04-02

**Type:** Reconstructed unit-testing prompt

> Act as a frontend unit-test engineer. Use the code produced by P07-04-01. Write focused Vitest and React Testing Library tests where appropriate.
>
> 1. Unit test keypad interactions, expression length, clear, special keys, success/error/loading and no duplicate submissions using mocked API functions; do not call live Go services.
> 2. Test accessible labels, visible state and errors where applicable. Avoid fragile implementation-detail assertions.
> 3. Mock external modules, network access and timers when needed; these are **unit tests**, not backend integration tests or browser E2E tests.
> 4. Run `npm test -- --run` or the actual test script defined in `package.json`, plus TypeScript checks if available; report actual results only.
> 5. Do not alter application source merely to make tests pass without approval.

## Code Review Prompt — P07-04-03

**Type:** Reconstructed code-review prompt

> Act as a principal frontend engineer. Review P07-04-01 implementation and P07-04-02 unit tests against the approved  design and the existing PascalCalc source.
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
