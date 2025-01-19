-- Users Table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role TEXT NOT NULL CHECK (role IN ('admin', 'employee'))
);

-- Employees Table
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    position TEXT NOT NULL
);

-- Reviews Table
CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    employee_id INT NOT NULL REFERENCES employees(id) ON DELETE CASCADE,
    performance_review TEXT NOT NULL,
    comments TEXT[] DEFAULT ARRAY[]::TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Review Reviewers Table
CREATE TABLE review_reviewers (
    id SERIAL PRIMARY KEY,
    review_id INT NOT NULL REFERENCES reviews(id) ON DELETE CASCADE,
    reviewer_id INT NOT NULL REFERENCES employees(id) ON DELETE CASCADE,
    UNIQUE (review_id, reviewer_id)
);

-- Feedback Table
CREATE TABLE feedback (
    id SERIAL PRIMARY KEY,
    review_id INT NOT NULL REFERENCES reviews(id) ON DELETE CASCADE,
    submitted BOOLEAN DEFAULT FALSE
);