# Sezzle Calculator Take-Home

A full-stack calculator application built for a Senior Software Engineer take-home assignment. The project uses a Go backend with clean architecture, a React + TypeScript frontend, token-protected REST APIs, and focused unit tests on both sides.

## Project overview

The app exposes a single `POST /api/v1/calculate` endpoint and a responsive UI that supports:

- Addition
- Subtraction
- Multiplication
- Division
- Exponentiation
- Square root
- Percentage

The implementation intentionally balances engineering quality with take-home scope. The code aims to demonstrate maintainability, testability, security awareness, and pragmatic decision-making within a 2-4 hour exercise.

## Repository structure

```text
.
├── backend/
│   ├── cmd/
│   │   ├── api/
│   │   └── devtoken/
│   └── internal/
│       ├── adapters/
│       ├── application/
│       ├── domain/
│       ├── infrastructure/
│       └── ports/
├── frontend/
│   └── src/
├── .env.example
├── PROMPTS.md
└── README.md
```

## Architecture explanation

### Backend

The backend follows a lightweight clean architecture / hexagonal architecture:

- `domain/calculator`: pure business rules and operation strategies
- `application/calculate`: orchestration, request validation, and operation dispatch
- `ports`: abstractions for token validation and operation lookup
- `adapters/http`: REST handlers, DTOs, structured responses, and middleware
- `adapters/auth`: replaceable local JWT validator
- `infrastructure`: config, logging, and router composition
- `cmd/api`: application entrypoint

Patterns used pragmatically:

- Strategy pattern for calculator operations
- Registry/factory for supported operations
- Middleware chain for auth, CORS, logging, security headers, and recovery

### Frontend

The frontend uses a small but intentional React structure:

- `api/`: HTTP client and calculator API calls
- `components/`: reusable presentational components
- `hooks/`: calculator state and async submission logic
- `types/`: shared frontend API contracts
- `utils/`: operation metadata and validation rules

This keeps the UI simple while avoiding a monolithic component.

## Security decisions

The assignment asked for OAuth2-style Bearer token protection without pulling in a full identity provider. The backend uses a pragmatic local JWT middleware that:

- Requires `Authorization: Bearer <token>`
- Validates the token signature using `JWT_SECRET`
- Uses an interface-based validator so a real OIDC/JWKS implementation can replace it later

For evaluator convenience, the app also exposes a local demo token endpoint:

- `POST /api/v1/auth/token`
- accepts `client_id` and `client_secret`
- returns a short-lived Bearer token signed with the same local `JWT_SECRET`

This is intentionally a local/demo bootstrap flow so reviewers can exercise the protected API from the UI without manually generating JWTs at the terminal.

Additional hardening included:

- CORS restricted to the configured `FRONTEND_ORIGIN` value
- Basic security headers
- Request body size limit
- Safe JSON decoding with `DisallowUnknownFields`
- Generic internal error responses that do not leak implementation details
- No secrets hardcoded in source

## API contract

### Request

`POST /api/v1/calculate`

```json
{
  "operation": "add",
  "operands": [10, 5]
}
```

### Success response

```json
{
  "result": 15
}
```

### Error response

```json
{
  "error": {
    "code": "DIVISION_BY_ZERO",
    "message": "Division by zero is not allowed"
  }
}
```

Supported operations:

- `add`
- `subtract`
- `multiply`
- `divide`
- `power`
- `sqrt`
- `percentage`

## Environment variables

Use `.env.example` as a reference.

Backend:

- `SERVER_ADDRESS`: API bind address, default `:18080`
- `FRONTEND_ORIGIN`: allowed CORS origin, default `http://localhost:15173`
- `JWT_SECRET`: HMAC signing secret for local JWT validation
- `DEMO_CLIENT_ID`: local demo client identifier used by the token-issuing endpoint
- `DEMO_CLIENT_SECRET`: local demo client secret used by the token-issuing endpoint

Frontend:

- `VITE_API_BASE_URL`: backend base URL, default `http://localhost:18080`

## How to run the backend

1. Copy values from `.env.example` into your shell environment.
2. Start the API:

```bash
cd backend
export SERVER_ADDRESS=:18080
export FRONTEND_ORIGIN=http://localhost:15173
export JWT_SECRET='replace-with-a-long-random-secret'
export DEMO_CLIENT_ID='sezzle-demo-client'
export DEMO_CLIENT_SECRET='replace-with-a-demo-client-secret'
go run ./cmd/api
```

### Generate a local development token from the UI

Open the frontend and use:

- `Client ID`: the value from `DEMO_CLIENT_ID`
- `Client secret`: the value from `DEMO_CLIENT_SECRET`
- click `Generate token`

The UI will request a short-lived Bearer token from the backend and autofill the token field for calculator requests.

### Generate a local development token from the terminal

Use the helper command below after setting `JWT_SECRET`:

```bash
cd backend
go run ./cmd/devtoken
```

You can also customize the subject and expiration:

```bash
go run ./cmd/devtoken -sub sezzle-demo -expires-in 2h
```

## How to run the frontend

1. Set the frontend environment variables in your shell.
2. Start the Vite dev server.

```bash
cd frontend
export VITE_API_BASE_URL=http://localhost:18080
npm install
npm run dev
```

Open [http://localhost:15173](http://localhost:15173), enter the demo `Client ID` and `Client secret`, click `Generate token`, and then use the calculator.

## How to run with Docker

From the repository root:

```bash
cd /Users/jnsilvag/Documents/Codex/2026-05-01-you-are-helping-me-build-a/sezzle-calculator
export JWT_SECRET='replace-with-a-long-random-secret'
export DEMO_CLIENT_ID='sezzle-demo-client'
export DEMO_CLIENT_SECRET='replace-with-a-demo-client-secret'
docker compose up --build
```

This starts:

- Backend API on [http://localhost:18080](http://localhost:18080)
- Frontend UI on [http://localhost:15173](http://localhost:15173)

Generate a token from the running backend container:

```bash
docker compose exec backend sh -lc 'export JWT_SECRET="$JWT_SECRET"; devtoken'
```

For normal UI testing you do not need that terminal step. Instead, use these values directly in the frontend:

- `Client ID`: `sezzle-demo-client` or your configured `DEMO_CLIENT_ID`
- `Client secret`: the value of your configured `DEMO_CLIENT_SECRET`

Then click `Generate token` in the UI and use the calculator normally.

## API examples with curl

First generate a token:

```bash
cd backend
export JWT_SECRET='replace-with-a-long-random-secret'
TOKEN=$(go run ./cmd/devtoken)
```

Or request one from the demo token endpoint:

```bash
curl -X POST http://localhost:18080/api/v1/auth/token \
  -H "Content-Type: application/json" \
  -d '{"client_id":"sezzle-demo-client","client_secret":"replace-with-a-demo-client-secret"}'
```

Addition:

```bash
curl -X POST http://localhost:18080/api/v1/calculate \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"operation":"add","operands":[10,5]}'
```

Square root:

```bash
curl -X POST http://localhost:18080/api/v1/calculate \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"operation":"sqrt","operands":[81]}'
```

Division by zero:

```bash
curl -X POST http://localhost:18080/api/v1/calculate \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"operation":"divide","operands":[10,0]}'
```

## How to run tests and coverage

### Test organization

This project uses colocated unit tests as the default approach:

- Backend unit tests live next to the Go packages they verify using `*_test.go`
- Frontend unit tests live next to the React components and utilities they verify using `*.test.ts` and `*.test.tsx`

This is intentional and aligns with common Go and React project conventions because it keeps tests close to the implementation they cover.

Examples in this repository:

- Backend domain tests in `backend/internal/domain/calculator`
- Backend application tests in `backend/internal/application/calculate`
- Backend HTTP and middleware tests in `backend/internal/adapters`
- Frontend component tests in `frontend/src/components`
- Frontend utility tests in `frontend/src/utils`

There are also support or future-expansion test folders:

- `backend/tests/integration`
  This folder is reserved for future integration or end-to-end style backend tests if broader cross-layer scenarios are added later.
- `frontend/src/test`
  This folder currently contains shared frontend test setup such as Vitest configuration helpers.

Today, the active automated tests are mainly the colocated unit and handler tests described above.

### Backend tests

The backend test suite covers:

- Domain operation logic
- Request validation and edge cases
- HTTP handlers with `httptest`
- Auth and CORS middleware behavior

Run all backend tests:

```bash
cd backend
go test ./...
```

Because Go tests are colocated, `go test ./...` automatically discovers and runs all backend tests across the module, including:

- domain tests
- application service tests
- handler tests
- middleware tests
- auth validator tests

Run backend coverage:

```bash
cd backend
go test ./... -cover
```

### Frontend tests

The frontend test suite covers:

- Validation utilities
- Main app rendering
- Calculator form behavior and submission rules

Run all frontend tests:

```bash
cd frontend
npm test
```

Because frontend tests are colocated, Vitest automatically finds files such as:

- `src/App.test.tsx`
- `src/components/*.test.tsx`
- `src/utils/*.test.ts`

Run frontend coverage:

```bash
cd frontend
npm run coverage
```

### Frontend production build check

This is not a unit test, but it is useful to verify the TypeScript and production bundle are healthy:

```bash
cd frontend
npm run build
```

### Run all checks before submitting

If you want a simple pre-submit sequence:

```bash
cd backend
go test ./... -cover

cd ../frontend
npm test
npm run build
```

### Optional API smoke test with Docker running

If the Docker stack is already up, you can also verify the full request path manually:

```bash
curl -X POST http://localhost:18080/api/v1/calculate \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"operation":"add","operands":[10,5]}'
```

## Design trade-offs and assumptions

- The backend uses only the Go standard library for the API and JWT handling to keep dependencies minimal.
- The JWT validator supports local HS256 tokens only. That is sufficient for the assignment while preserving a clean seam for a future JWKS-backed implementation.
- The API intentionally exposes a single endpoint because the assignment centered on operation dispatch rather than broader resource modeling.
- The frontend assumes a token is provisioned out-of-band through environment configuration. No secrets are stored in browser code beyond what a local developer chooses to inject for testing.
- Logging is intentionally simple and JSON-based rather than introducing a larger observability stack.

## What I would improve in production

- Replace local HMAC JWT validation with real OIDC discovery + JWKS rotation
- Add request IDs and richer structured logging
- Add integration tests covering full backend request flows
- Add rate limiting and possibly audit logging for protected endpoints
- Add better numeric formatting and accessibility refinements in the frontend
- Add containerization and deployment manifests
- Add CI for linting, tests, coverage thresholds, and dependency scanning

## Notes on dependency audit output

`npm audit` currently reports moderate transitive vulnerabilities from the frontend dependency tree. I did not run `npm audit fix --force` because it can introduce breaking upgrades during a take-home. In a production repo, I would review and address those upgrades deliberately.
