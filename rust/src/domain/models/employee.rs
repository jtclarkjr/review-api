use serde::{Deserialize, Serialize};
use sqlx::FromRow;
use utoipa::ToSchema;

#[derive(Debug, FromRow, Serialize, Deserialize, ToSchema)]
pub struct Employee {
    pub id: i32,
    pub email: String,
    pub position: String,
}

#[derive(Debug, Deserialize, ToSchema)]
pub struct CreateEmployeeRequest {
    pub email: String,
    pub position: String,
}

#[derive(Debug, Deserialize, ToSchema)]
pub struct UpdateEmployeeRequest {
    pub email: Option<String>,
    pub position: Option<String>,
}
