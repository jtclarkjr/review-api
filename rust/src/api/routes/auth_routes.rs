use axum::{extract::State, Json};
use sqlx::PgPool;

use crate::application::AuthService;
use crate::domain::models::{LoginRequest, LoginResponse};
use crate::infrastructure::Result;

pub async fn login(
    State(pool): State<PgPool>,
    Json(req): Json<LoginRequest>,
) -> Result<Json<LoginResponse>> {
    let response = AuthService::login(&pool, &req).await?;
    Ok(Json(response))
}
