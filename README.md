# Yelp Clone

A simplified clone of Yelp built with Go, Bun, and PostgreSQL.  
Designed to manage users, businesses, menus, and reviews with a RESTful API.

## Features
- User authentication and profiles
- Businesses with categories, amenities, images, hours
- Menus, menu items, and menu categories
- Reviews with ratings and images
- Collections and business following
- Full REST API for frontend consumption

## Tech Stack
- **Backend:** Go, Bun ORM, PostgreSQL
- **Database:** PostgreSQL
- **Migrations:** Bun migrate
- **Testing:** TBD
- **Environment management:** `.env` files, `APP_ENV` variable

## Getting Started

## Prerequisites
- Go 1.21+
- PostgreSQL 15+
- `make` (optional)

## Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/promingy/yelp-clone.git
   cd yelp-clone/backend
   ```
2. Create `.env` file in `/backend`:
   ```env
   APP_ENV=dev
   DATABASE_URL=postgres://username:password@localhost:5432/yelp_clone?sslmode=disable
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Run database migrations:
   ```bash
   go run ./cmd/migrate
   ```
5. Start the backend server:
   ```bash
   go run ./cmd/api
   ```

## Running in Production
- Set `APP_ENV=prod`
- Ensure `DATABASE_URL` points to production DB
- Run migrations and start the server as above

## API Documentation
See [API.md](API.md) for endpoint details.

<!-- ## Contributing
See [CONTRIBUTIONS.md](CONTRIBUTIONS.md) for guidelines. -->
