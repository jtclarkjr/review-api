# Employee Review REST API - Rust

A high-performance REST API for managing employee reviews, built with Rust and Axum. This implementation provides separate functionality for admin and employee roles with JWT authentication.

## Features

### Admin View
Admins can:
- **Manage Employees**:
    - Add new employees
    - Remove employees
    - Update employee information
    - View the list of employees

- **Manage Performance Reviews**:
    - Add new performance reviews
    - Update existing performance reviews
    - View all performance reviews

- **Assign Participants**:
    - Assign employees to provide feedback for another employee's performance review

### Employee View
Employees can:
- **View Assigned Performance Reviews**:
    - List performance reviews that require their feedback

- **Submit Feedback**:
    - Provide feedback for assigned performance reviews

## API Endpoints

### Authentication
- **Login**  
  `POST /login`  
  Authenticate and receive a JWT token

### Admin Endpoints

#### Employee Management
- **Add Employee**  
  `POST /admin/employees`  
  Add a new employee

- **Update Employee**  
  `PUT /admin/employees/{id}`  
  Update an existing employee's information

- **Remove Employee**  
  `DELETE /admin/employees/{id}`  
  Remove an employee by ID

- **View Employees**  
  `GET /admin/employees`  
  Retrieve a list of all employees

#### Performance Review Management
- **Add Performance Review**  
  `POST /admin/reviews`  
  Create a new performance review

- **Update Performance Review**  
  `PUT /admin/reviews/{id}`  
  Update details of an existing performance review

- **View Performance Reviews**  
  `GET /admin/reviews`  
  Retrieve all performance reviews

#### Assign Participants
- **Assign Reviewer to Performance Review**  
  `POST /admin/reviews/{review_id}/assign`  
  Assign an employee as a reviewer for a specific performance review

### Employee Endpoints

#### Performance Reviews
- **List Assigned Reviews**  
  `GET /employee/reviews`  
  Retrieve a list of performance reviews assigned to the employee that require feedback

- **Submit Feedback**  
  `POST /employee/reviews/{review_id}/feedback`  
  Submit feedback for an assigned performance review

## Technology Stack

- **Axum** - Modern web framework built on Tokio
- **SQLx** - Async PostgreSQL driver with compile-time query verification
- **Tower** - Middleware and service layers
- **JWT** - Token-based authentication
- **Bcrypt** - Password hashing
- **Tokio** - Async runtime

## Prerequisites

- **Rust** (version 1.75 or higher with edition 2024 support)
- **PostgreSQL** database

## Installation

1. Clone the repository:
   ```bash
   cd rust
   ```

2. Set up environment variables:
   ```bash
   cp .env.example .env
   ```

3. Configure your `.env` file:
   ```bash
   DATABASE_URL=postgresql://user:password@localhost:5432/review_api
   JWT_SECRET=your-secret-key-here
   ```

4. Run the database schema (from the parent go directory):
   ```bash
   psql $DATABASE_URL < ../go/db/schema.sql
   ```

5. Build and run:
   ```bash
   cargo run
   ```

The server will start on `http://localhost:8080`

## Development

### Running in development mode
```bash
cargo run
```

### Building for production
```bash
cargo build --release
```

### Running with logging
```bash
RUST_LOG=debug cargo run
```

## Usage Examples

### Login
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "admin"
  }'
```

### Add Employee (Admin)
```bash
curl -X POST http://localhost:8080/admin/employees \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "position": "Software Engineer"
  }'
```

### Get Assigned Reviews (Employee)
```bash
curl -X GET http://localhost:8080/employee/reviews \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Submit Feedback (Employee)
```bash
curl -X POST http://localhost:8080/employee/reviews/1/feedback \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "comment": "Excellent performance throughout the quarter"
  }'
```

## Architecture

This implementation follows clean architecture and domain-driven design principles with Rust best practices:

- **Layered Architecture**: Clear separation between API, Application, Domain, and Infrastructure layers
- **Type Safety**: Leverages Rust's strong type system with compile-time query verification via SQLx
- **Error Handling**: Custom error types with proper propagation using the `?` operator
- **Async/Await**: Fully asynchronous using Tokio runtime
- **Repository Pattern**: Database operations abstracted into repository layer
- **Service Layer**: Business logic encapsulated in service layer
- **JWT Authentication**: Implemented as a custom Axum extractor
- **Dependency Injection**: State management through Axum's extractors

## Project Structure

```
src/
├── main.rs                    - Application entry point
├── lib.rs                     - Library root
├── api/
│   ├── mod.rs                - API layer module
│   ├── router.rs             - Route configuration
│   └── routes/
│       ├── mod.rs            - Route handlers module
│       ├── auth_routes.rs    - Authentication endpoints
│       ├── admin_routes.rs   - Admin endpoints
│       └── employee_routes.rs - Employee endpoints
├── application/
│   ├── mod.rs                - Application layer module
│   └── services/
│       ├── mod.rs            - Services module
│       ├── auth_service.rs   - Authentication business logic
│       ├── employee_service.rs - Employee business logic
│       └── review_service.rs - Review business logic
├── domain/
│   ├── mod.rs                - Domain layer module
│   └── models/
│       ├── mod.rs            - Models module
│       ├── user.rs           - User domain models
│       ├── employee.rs       - Employee domain models
│       ├── review.rs         - Review domain models
│       └── common.rs         - Common models
└── infrastructure/
    ├── mod.rs                - Infrastructure layer module
    ├── auth.rs               - JWT authentication utilities
    ├── error.rs              - Error handling
    ├── database/
    │   ├── mod.rs            - Database module
    │   ├── pool.rs           - Connection pooling
    │   ├── user_repository.rs - User data access
    │   ├── employee_repository.rs - Employee data access
    │   └── review_repository.rs - Review data access
    └── config/
        ├── mod.rs            - Config module
        └── settings.rs       - Application settings
```

### Layer Responsibilities

**Domain Layer**: Pure business entities and DTOs, no external dependencies

**Application Layer**: Business logic and use cases, orchestrates domain objects and repositories

**Infrastructure Layer**: External concerns like database, authentication, configuration

**API Layer**: HTTP handlers, request/response mapping, routing
