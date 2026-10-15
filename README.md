# URL Shortener (Go + PostgreSQL + Redis)

A backend URL shortener built from scratch in Go, demonstrating REST API design, secure authentication, relational database persistence, caching, and containerization.

## Features
- **REST API** — create short links and redirect to original URLs
- **Authentication** — user registration and login with bcrypt password hashing
- **Session management** — token-based auth protecting private routes
- **Persistent storage** — PostgreSQL for users, URLs, and sessions
- **Caching** — Redis cache-aside pattern for fast redirects (cache hit/miss logic)
- **Containerized** — multi-stage Docker build for a lightweight, portable deployment

## Tech Stack
- **Language:** Go (standard library `net/http`, no framework)
- **Database:** PostgreSQL
- **Cache:** Redis
- **Auth:** bcrypt password hashing, token-based sessions
- **Containerization:** Docker (multi-stage build)

## API Endpoints

| Method | Endpoint | Description | Auth required |
|--------|----------|--------------|----------------|
| POST | `/register` | Create a new user account | No |
| POST | `/login` | Log in, returns a session token | No |
| POST | `/shorten` | Create a short link from a long URL | No |
| GET | `/{shortCode}` | Redirect to the original URL | No |
| GET | `/my-urls` | Protected example route | Yes (Authorization header) |

## Architecture Notes
- Passwords are never stored in plain text — only bcrypt hashes
- Session tokens are random 32-character strings, stored server-side in a `sessions` table
- Redirects check Redis first (cache-aside pattern); on a cache miss, the app queries PostgreSQL and populates the cache for subsequent requests (1-hour TTL)
- The Docker image uses a multi-stage build — a full Go environment compiles the binary, then the final image only contains the compiled executable on a minimal Alpine base

## Running Locally

### Prerequisites
- Go 1.26+
- PostgreSQL running locally
- Redis (via Docker: `docker run -d -p 6379:6379 redis`)

### Setup
1. Clone the repo
2. Create a PostgreSQL database and run the schema (see `schema.sql`)
3. Update the connection strings in `main.go` with your credentials
4. Run:
```bash
go run main.go

## Cloud Deployment

This project has been successfully deployed to a cloud environment, validating the full stack beyond local development:

- **Database:** [Supabase](https://supabase.com) (managed PostgreSQL)
- **Cache:** [Upstash](https://upstash.com) (managed Redis)
- **Compute:** [Back4App Containers](https://www.back4app.com) (Docker-based hosting)
- **Configuration:** all database/cache connections are read from environment variables (`DATABASE_URL`, `REDIS_URL`), not hardcoded — standard practice for deployable services

### A real deployment bug I hit and fixed
After deploying, the app failed to connect to PostgreSQL with `dial tcp ... connect: cannot assign requested address`. The cause: Supabase's **direct connection** endpoint defaults to IPv6, which the hosting platform's network couldn't route to. The fix was switching to Supabase's **connection pooler** endpoint, which supports IPv4 — a good reminder that cloud networking issues are often about *which* endpoint you connect to, not whether your code is correct.

*Note: the live deployment uses a free-tier hosting plan intended for demos, so it is not kept permanently running. The steps above are fully reproducible — clone the repo, provision your own free Supabase/Upstash instances, and deploy.*