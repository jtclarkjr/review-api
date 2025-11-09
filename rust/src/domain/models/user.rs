use serde::{Deserialize, Serialize};
use sqlx::FromRow;
use utoipa::ToSchema;

#[derive(Debug, Clone, Serialize, Deserialize, sqlx::Type, PartialEq, ToSchema)]
#[sqlx(type_name = "text")]
#[serde(rename_all = "lowercase")]
pub enum Role {
    #[sqlx(rename = "admin")]
    Admin,
    #[sqlx(rename = "employee")]
    Employee,
}

#[derive(Debug, FromRow, Serialize, Deserialize, ToSchema)]
pub struct User {
    pub id: i32,
    pub email: String,
    #[serde(skip_serializing)]
    pub password: String,
    pub role: Role,
}

#[derive(Debug, Deserialize, ToSchema)]
pub struct LoginRequest {
    pub email: String,
    pub password: String,
}

#[derive(Debug, Serialize, ToSchema)]
pub struct LoginResponse {
    pub token: String,
}
