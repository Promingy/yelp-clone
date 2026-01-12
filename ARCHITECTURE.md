# Architecture Overview

This document explains the structure and design of the Yelp Clone backend.

## Folder Structure
```
/backend
├── cmd/
│   ├── api/          # Main server entrypoint
│   └── migrate/      # Database migrations
├── internal/
│   ├── db/           # Database connection logic
│   ├── models/       # Bun models for all tables
│   ├── migrations/   # Go-based migration files
│   ├── routes/       # API route setup
│   └── handlers/     # HTTP handlers
├── pkg/              # Optional reusable packages
├── .env              # Environment variables
└── go.mod
```
## Key Components

### Database
- **Bun ORM** for schema management and query building
- **PostgreSQL** as the primary database
- **Models** in `internal/models` represent tables such as:
  - `User`, `Business`, `Category`, `Menu`, `MenuItem`, `Review`, etc.
- **Migrations** in `internal/migrations` manage schema evolution.

### Server
- **Entry point:** `cmd/api/main.go`
- **Routing:** `bunrouter` with handlers in `internal/handlers`
- **Middleware:** Logging, error handling, authentication (planned)

### Data Flow
1. Client makes an HTTP request to `/api`.
2. Request is routed via `bunrouter` to a handler.
3. Handler interacts with `bun` models and the PostgreSQL database.
4. Response is serialized (JSON) and returned to the client.

### Environment Management
- `.env` file for local dev
- `APP_ENV` variable toggles `dev` vs `prod`
- `DATABASE_URL` contains connection info

---

## Conventions
- RESTful routes under `/api`
- Timestamps: `created_at`, `updated_at` (auto-populated)
- Join tables use composite primary keys