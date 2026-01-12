# Project Conventions

This document outlines coding, naming, and architectural conventions used in the Yelp Clone backend.

## Folder & Package Conventions
- **`cmd/`**: Entry points (`api`, `migrate`)
- **`internal/`**: Private code for this service
  - `db/`: Database connection
  - `models/`: Bun models
  - `migrations/`: Go migrations
  - `handlers/`: HTTP handlers
  - `routes/`: API routing
- **`pkg/`**: Optional reusable packages

## Database Conventions
- Tables: `snake_case`, plural (e.g., `businesses`, `users`)
- Columns: `snake_case` (e.g., `created_at`, `is_verified`)
- Join tables: `table1_table2` with composite PKs
- Timestamps: `created_at`, `updated_at` with defaults
- Foreign keys: `{related_table}_id`

## Go Conventions
- **Structs:** PascalCase (exported), with Bun tags for DB mapping
- **Struct fields for DB columns:** Tagged with `bun:"column_name[,options]"`
- **Booleans:** Default false unless otherwise specified
- **Time:** `time.Time` for timestamps, `*time.Time` if nullable

## API & Routing
- RESTful endpoints, grouped by resource:
  - `/users`
  - `/businesses`
  - `/menus`
  - `/reviews`
- Handlers return JSON
- Error handling via `bunrouter.HandlerFunc` returning `error`

## Naming Conventions
- Files: `snake_case.go` (except `main.go`)
- Packages: `camelcase`
- Join tables: `table1_table2`
- Enums or constants: PascalCase

## Miscellaneous
- Use `.env` for environment variables
- `APP_ENV` values: `dev` or `prod`
- Git commit messages follow conventional commits style: `feat:`, `fix:`, `chore:`