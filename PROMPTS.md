# AI Prompts Used During Development

This file documents the main prompts used to guide AI-assisted development for the assignment.

## Initial architecture prompt

```text
You are helping me build a take-home assignment for a Senior Software Engineer role at Sezzle.

The assignment is a full-stack calculator app with:
- Backend: Go
- Frontend: React with TypeScript
- REST API
- Unit tests
- README with setup instructions, API examples, design decisions, and AI prompts used

Important goal:
The solution must look like it was built by a Senior Software Engineer: clean architecture, maintainability, security awareness, testability, idiomatic Go, idiomatic React, and pragmatic scope for a 2–4 hour take-home assignment.

Please generate the project step by step, not all at once. Before coding, propose the architecture and file structure.

Functional requirements:
Backend calculator operations:
- Addition
- Subtraction
- Multiplication
- Division
- Optional but preferred: exponentiation, square root, percentage

Frontend:
- Intuitive calculator UI
- Input validation
- Error handling
- Responsive layout
- TypeScript types
- Clean component structure

Backend:
- REST API endpoints for calculator operations
- Validate input
- Handle edge cases such as division by zero, invalid numbers, square root of negative numbers
- Return JSON responses
- Include structured error responses

Security requirements:
- Protect API endpoints using OAuth2-style Bearer token validation.
- Since this is a take-home app, implement a pragmatic local JWT/OAuth2-compatible middleware:
  - Require Authorization: Bearer <token>
  - Validate token signature using a local secret or JWKS-like abstraction
  - Make the auth layer replaceable by a real OAuth2/OIDC provider later
- Do not hardcode secrets in source code; use environment variables.
- Add CORS configuration restricted to the frontend origin.
- Add basic security headers.
- Add request size limits and safe JSON decoding.
- Make sure errors do not leak internal implementation details.

Architecture requirements:
Backend should use hexagonal architecture / clean architecture:
- domain: calculator business logic
- application/usecases: operation orchestration
- ports: interfaces
- adapters/http: REST handlers
- adapters/auth: token validation middleware
- infrastructure/config/logging
- cmd/api/main.go

Use SOLID principles:
- Business logic must be independent from HTTP
- Handlers should be thin
- Validation should be clear and testable
- Auth should be abstracted behind an interface
- Avoid unnecessary overengineering

Design patterns to apply pragmatically:
- Strategy pattern for calculator operations
- Factory or registry for supported operations
- Middleware chain for cross-cutting concerns such as auth, logging, CORS, and recovery

Testing:
Backend:
- Unit tests for calculator operations
- Unit tests for validation and edge cases
- Handler tests using httptest
- Auth middleware tests
- Include coverage command in README

Frontend:
- Unit tests for components and validation
- Use Vitest and React Testing Library
- Test success and error scenarios

Frontend architecture:
- React + TypeScript + Vite
- API client layer separated from UI components
- Reusable components
- Form/input validation
- Environment variable for backend URL
- Clear error and loading states
- Avoid storing secrets in frontend

API design:
Use one main endpoint:
POST /api/v1/calculate

Request:
{
  "operation": "add",
  "operands": [10, 5]
}

Response:
{
  "result": 15
}

Error response:
{
  "error": {
    "code": "DIVISION_BY_ZERO",
    "message": "Division by zero is not allowed"
  }
}

Supported operations:
add, subtract, multiply, divide, power, sqrt, percentage

Deliverables:
- backend/
- frontend/
- docker-compose.yml optional
- README.md
- PROMPTS.md containing the prompts used during development

README must include:
- Project overview
- Architecture explanation
- Security decisions
- API examples with curl
- How to run backend
- How to run frontend
- How to run tests and coverage
- Environment variables
- Design trade-offs and assumptions
- What would be improved in production

Do not use heavy frameworks unless justified.
For Go, prefer standard library plus minimal dependencies.
For React, use Vite, TypeScript, Vitest, and React Testing Library.
```

## Implementation prompts

```text
Start by proposing:
1. Repository structure
2. Backend architecture
3. Frontend architecture
4. API contract
5. Security approach
6. Testing strategy
7. Implementation plan
```

```text
Looks good, go ahead.
```

```text
Looks good, go ahead with the next step.
```

```text
Let's continue with the next step.
```

## How AI assistance was used

AI assistance was used for:

- Initial architecture proposal
- Scaffolding backend and frontend structure
- Implementing tests
- Drafting README and prompt documentation

All generated code and documentation were reviewed and iterated on during development.
