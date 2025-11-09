use axum::{extract::{Path, State}, Json};
use sqlx::PgPool;

use crate::application::{EmployeeService, ReviewService};
use crate::domain::models::*;
use crate::infrastructure::{require_admin, AuthUser, Result};

pub async fn create_employee(
    State(pool): State<PgPool>,
    user: AuthUser,
    Json(req): Json<CreateEmployeeRequest>,
) -> Result<Json<Employee>> {
    require_admin(&user)?;
    let employee = EmployeeService::create(&pool, &req).await?;
    Ok(Json(employee))
}

pub async fn get_employees(
    State(pool): State<PgPool>,
    user: AuthUser,
) -> Result<Json<Vec<Employee>>> {
    require_admin(&user)?;
    let employees = EmployeeService::get_all(&pool).await?;
    Ok(Json(employees))
}

pub async fn update_employee(
    State(pool): State<PgPool>,
    user: AuthUser,
    Path(id): Path<i32>,
    Json(req): Json<UpdateEmployeeRequest>,
) -> Result<Json<Employee>> {
    require_admin(&user)?;
    let employee = EmployeeService::update(&pool, id, &req).await?;
    Ok(Json(employee))
}

pub async fn delete_employee(
    State(pool): State<PgPool>,
    user: AuthUser,
    Path(id): Path<i32>,
) -> Result<Json<serde_json::Value>> {
    require_admin(&user)?;
    EmployeeService::delete(&pool, id).await?;
    Ok(Json(serde_json::json!({"message": "Employee deleted"})))
}

pub async fn create_review(
    State(pool): State<PgPool>,
    user: AuthUser,
    Json(req): Json<CreateReviewRequest>,
) -> Result<Json<Review>> {
    require_admin(&user)?;
    let review = ReviewService::create(&pool, &req).await?;
    Ok(Json(review))
}

pub async fn get_reviews(
    State(pool): State<PgPool>,
    user: AuthUser,
) -> Result<Json<Vec<Review>>> {
    require_admin(&user)?;
    let reviews = ReviewService::get_all(&pool).await?;
    Ok(Json(reviews))
}

pub async fn update_review(
    State(pool): State<PgPool>,
    user: AuthUser,
    Path(id): Path<i32>,
    Json(req): Json<UpdateReviewRequest>,
) -> Result<Json<Review>> {
    require_admin(&user)?;
    let review = ReviewService::update(&pool, id, &req).await?;
    Ok(Json(review))
}

pub async fn assign_reviewer(
    State(pool): State<PgPool>,
    user: AuthUser,
    Path(review_id): Path<i32>,
    Json(req): Json<AssignReviewerRequest>,
) -> Result<Json<ReviewReviewer>> {
    require_admin(&user)?;
    let reviewer = ReviewService::assign_reviewer(&pool, review_id, &req).await?;
    Ok(Json(reviewer))
}
