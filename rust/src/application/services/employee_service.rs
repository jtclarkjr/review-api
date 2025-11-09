use sqlx::PgPool;

use crate::domain::models::{CreateEmployeeRequest, Employee, UpdateEmployeeRequest};
use crate::infrastructure::{AppError, EmployeeRepository, Result};

pub struct EmployeeService;

impl EmployeeService {
    pub async fn create(pool: &PgPool, req: &CreateEmployeeRequest) -> Result<Employee> {
        EmployeeRepository::create(pool, req).await
    }

    pub async fn get_all(pool: &PgPool) -> Result<Vec<Employee>> {
        EmployeeRepository::find_all(pool).await
    }

    pub async fn get_by_id(pool: &PgPool, id: i32) -> Result<Employee> {
        EmployeeRepository::find_by_id(pool, id).await
    }

    pub async fn get_by_email(pool: &PgPool, email: &str) -> Result<Employee> {
        EmployeeRepository::find_by_email(pool, email)
            .await?
            .ok_or_else(|| AppError::NotFound("Employee not found".to_string()))
    }

    pub async fn update(pool: &PgPool, id: i32, req: &UpdateEmployeeRequest) -> Result<Employee> {
        EmployeeRepository::update(pool, id, req).await
    }

    pub async fn delete(pool: &PgPool, id: i32) -> Result<()> {
        let deleted = EmployeeRepository::delete(pool, id).await?;
        if !deleted {
            return Err(AppError::NotFound("Employee not found".to_string()));
        }
        Ok(())
    }
}
