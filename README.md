# Review API

An employee review management API with dual implementations in Go and TypeScript/NestJS.

## Project Structure

- **`go/`** - Go implementation using a custom router with raw SQL queries
- **`ts/`** - TypeScript/NestJS implementation with Prisma ORM

## Features

- Employee management
- Authentication & authorization (JWT)
- Admin functionality
- API documentation (Swagger/OpenAPI)
- Database integration (PostgreSQL)

## Getting Started

See the README files in each implementation directory (`go/` and `ts/`) for setup and installation instructions.

## Technologies Used

### Go
- Custom router (`github.com/jtclarkjr/router-go`)
- Raw SQL queries with PostgreSQL driver
- JWT authentication
- Swagger documentation

### TypeScript/NestJS
- NestJS framework
- Prisma ORM for database abstraction
- JWT authentication with Passport
- Swagger documentation
- Bun runtime
