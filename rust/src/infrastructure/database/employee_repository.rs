use sqlx::PgPool;

use crate::domain::models::{CreateEmployeeRequest, Employee, UpdateEmployeeRequest};
use crate::infrastructure::error::{AppError, Result};

pub struct EmployeeRepository;

impl EmployeeRepository {
    pub async fn create(pool: &PgPool, req: &CreateEmployeeRequest) -> Result<Employee> {
        let employee = sqlx::query_as::<_, Employee>(
            "INSERT INTO employees (email, position) VALUES ($1, $2) RETURNING *",
        )
        .bind(&req.email)
        .bind(&req.position)
        .fetch_one(pool)
        .await?;

        Ok(employee)
    }

    pub async fn find_all(pool: &PgPool) -> Result<Vec<Employee>> {
        let employees = sqlx::query_as::<_, Employee>("SELECT * FROM employees")
            .fetch_all(pool)
            .await?;

        Ok(employees)
    }

    pub async fn find_by_id(pool: &PgPool, id: i32) -> Result<Employee> {
        let employee = sqlx::query_as::<_, Employee>("SELECT * FROM employees WHERE id = $1")
            .bind(id)
            .fetch_one(pool)
            .await?;

        Ok(employee)
    }

    pub async fn find_by_email(pool: &PgPool, email: &str) -> Result<Option<Employee>> {
        let employee = sqlx::query_as::<_, Employee>("SELECT * FROM employees WHERE email = $1")
            .bind(email)
            .fetch_optional(pool)
            .await?;

        Ok(employee)
    }

    pub async fn update(pool: &PgPool, id: i32, req: &UpdateEmployeeRequest) -> Result<Employee> {
        let mut query = String::from("UPDATE employees SET ");
        let mut updates = Vec::new();
        let mut bind_count = 1;

        if req.email.is_some() {
            updates.push(format!("email = ${}", bind_count));
            bind_count += 1;
        }
        if req.position.is_some() {
            updates.push(format!("position = ${}", bind_count));
            bind_count += 1;
        }

        if updates.is_empty() {
            return Err(AppError::BadRequest("No fields to update".to_string()));
        }

        query.push_str(&updates.join(", "));
        query.push_str(&format!(" WHERE id = ${} RETURNING *", bind_count));

        let mut q = sqlx::query_as::<_, Employee>(&query);

        if let Some(ref email) = req.email {
            q = q.bind(email);
        }
        if let Some(ref position) = req.position {
            q = q.bind(position);
        }
        q = q.bind(id);

        let employee = q.fetch_one(pool).await?;

        Ok(employee)
    }

    pub async fn delete(pool: &PgPool, id: i32) -> Result<bool> {
        let result = sqlx::query("DELETE FROM employees WHERE id = $1")
            .bind(id)
            .execute(pool)
            .await?;

        Ok(result.rows_affected() > 0)
    }
}
