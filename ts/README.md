# Employee Review REST API (TypeScript/NestJS)

<p align="center">
  <a href="http://nestjs.com/" target="blank"><img src="https://nestjs.com/img/logo-small.svg" width="120" alt="Nest Logo" /></a>
</p>

This project is a **TypeScript/NestJS-based REST API** for managing employee reviews, built using the NestJS framework with Prisma ORM. It provides separate functionality for **admin** and **employee** roles. Admins can manage employees, performance reviews, and assignments, while employees can view and submit feedback for assigned performance reviews.

## Features

### Admin View
Admins can:
- **Manage Employees**:
    - Add new employees.
    - Remove employees.
    - Update employee information.
    - View the list of employees.

- **Manage Performance Reviews**:
    - Add new performance reviews.
    - Update existing performance reviews.
    - View all performance reviews.

- **Assign Participants**:
    - Assign employees to provide feedback for another employee's performance review.

### Employee View
Employees can:
- **View Assigned Performance Reviews**:
    - List performance reviews that require their feedback.

- **Submit Feedback**:
    - Provide feedback for assigned performance reviews.

---

## API Endpoints

### Authentication
- **Login**  
  `POST /auth/login`  
  Authenticate user and receive JWT token.

### Admin Endpoints (Protected)

#### Employees Management
- **Add Employee**  
  `POST /admin/employees`  
  Add a new employee.

- **Update Employee**  
  `PUT /admin/employees/{id}`  
  Update an existing employee's information.

- **Remove Employee**  
  `DELETE /admin/employees/{id}`  
  Remove an employee by ID.

- **View Employees**  
  `GET /admin/employees`  
  Retrieve a list of all employees.

#### Performance Reviews Management
- **Add Performance Review**  
  `POST /admin/reviews`  
  Create a new performance review.

- **Update Performance Review**  
  `PUT /admin/reviews/{id}`  
  Update details of an existing performance review.

- **View Performance Reviews**  
  `GET /admin/reviews`  
  Retrieve all performance reviews.

---

### Employee Endpoints (Protected)

#### Performance Reviews
- **List Assigned Reviews**  
  `GET /feedback`  
  Retrieve a list of performance reviews assigned to the employee that require feedback.

- **Submit Feedback**  
  `POST /feedback`  
  Submit feedback for an assigned performance review.

- **Get My Feedbacks**  
  `GET /feedback/me`  
  Get feedbacks submitted by the current user.

- **Get Review Feedbacks**  
  `GET /feedback/review/{reviewId}`  
  Get feedbacks for a specific review.

---

## Tech Stack

- **Framework**: [NestJS](https://nestjs.com/) - A progressive Node.js framework
- **Language**: TypeScript
- **ORM**: [Prisma](https://www.prisma.io/)
- **Database**: PostgreSQL
- **Authentication**: JWT with Passport
- **Documentation**: Swagger/OpenAPI
- **Runtime**: Bun (can also use Node.js)

## Requirements

### Prerequisites
- **Bun** (or Node.js 18+ with npm/yarn)
- **PostgreSQL** as the database backend

### Dependencies
- **@nestjs/core** - NestJS framework core
- **@prisma/client** - Prisma ORM client
- **@nestjs/passport & @nestjs/jwt** - Authentication
- **bcrypt** - Password hashing
- **class-validator & class-transformer** - DTO validation

---

## Installation

1. Clone the repository and install dependencies:
```bash
git clone <repository-url>
cd review-api/ts
bun install
```

2. Set up the database:
```bash
# Generate Prisma client
bunx prisma generate

# Run database migrations
bunx prisma migrate dev
```

3. Set up environment variables:
```bash
# Copy environment template
cp .env.example .env

# Update .env with your database URL and JWT secret
DATABASE_URL="postgresql://username:password@localhost:5432/review_db"
JWT_SECRET="your-secret-key"
```

## Running the app

```bash
# development
$ bun start

# watch mode
$ bun start:dev

# production mode
$ bun start:prod
```

## Database Management

```bash
# Open Prisma Studio (database GUI)
$ bunx prisma studio

# Reset database
$ bunx prisma migrate reset

# Deploy migrations to production
$ bunx prisma migrate deploy
```

## API Documentation

This project uses @nestjs/swagger to automatically generate OpenAPI documentation.

### Running Swagger Locally

1. **Start the NestJS application** in development mode:
```bash
# Development with auto-reload
bun run start:dev

# Or regular development mode
bun run start
```

2. **Wait for the application to fully start** - you should see:
```bash
[Nest] LOG [NestApplication] Nest application successfully started +X ms
```

3. **Access the Swagger UI** in your browser at:
```
http://localhost:3000/api
```

4. **Explore the API documentation**:
   - Browse all available endpoints organized by tags
   - Test endpoints directly from the browser
   - View request/response schemas
   - Use the "Authorize" button to add your JWT token for protected endpoints

### Authentication in Swagger

For protected endpoints (those with a lock icon):
1. First, use the `/auth/login` endpoint to get a JWT token
2. Click the "Authorize" button at the top of the Swagger page
3. Enter your token in the format: `Bearer your-jwt-token-here`
4. Now you can test protected endpoints

### How Swagger Documentation Works

The API documentation is generated automatically from:
- Controller decorators (`@Controller()`, `@Get()`, `@Post()`, etc.)
- DTO classes with `@ApiProperty()` decorators
- Additional Swagger decorators:
  - `@ApiTags()` - Groups endpoints by controller
  - `@ApiOperation()` - Describes what an endpoint does
  - `@ApiResponse()` - Documents response types and status codes
  - `@ApiParam()` - Documents path parameters
  - `@ApiBody()` - Documents request body schema
  - `@ApiBearerAuth()` - Indicates JWT authentication

The Swagger configuration is set up in `main.ts` using `SwaggerModule.setup()`.

## Local Development

### Sample API Requests

#### Login
```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"email": "admin@company.com", "password": "adminpassword"}' \
http://localhost:3000/auth/login
```

#### Add Employee (Admin)
```bash
curl -X POST http://localhost:3000/admin/employees \
-H "Authorization: Bearer <your-jwt-token>" \
-H "Content-Type: application/json" \
-d '{"name": "John Doe", "email": "john@example.com", "position": "Developer"}'
```

#### Submit Feedback (Employee)
```bash
curl -X POST http://localhost:3000/feedback \
-H "Authorization: Bearer <your-jwt-token>" \
-H "Content-Type: application/json" \
-d '{"reviewId": "review-id", "feedback": "Great work on the project!"}'
```

## Test

```bash
# unit tests
$ bun test

# e2e tests
$ bun test:e2e

# test coverage
$ bun test:cov
```

## NestJS Resources

- **NestJS Documentation**: [https://docs.nestjs.com](https://docs.nestjs.com)
- **Prisma Documentation**: [https://www.prisma.io/docs](https://www.prisma.io/docs)
- **Swagger/OpenAPI**: [https://swagger.io/](https://swagger.io/)

## Stay in touch

- **NestJS Author** - [Kamil Myśliwiec](https://kamilmysliwiec.com)
- **Website** - [https://nestjs.com](https://nestjs.com/)
- **Twitter** - [@nestframework](https://twitter.com/nestframework)

## License

Nest is [MIT licensed](LICENSE).
