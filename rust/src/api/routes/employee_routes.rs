use axum::{extract::{Path, State}, Json};
use sqlx::PgPool;

use crate::application::ReviewService;
use crate::domain::models::{Review, SubmitFeedbackRequest};
use crate::infrastructure::{AuthUser, Result};

pub async fn get_assigned_reviews(
    State(pool): State<PgPool>,
    user: AuthUser,
) -> Result<Json<Vec<Review>>> {
    let reviews = ReviewService::get_assigned_reviews(&pool, &user.email).await?;
    Ok(Json(reviews))
}

pub async fn submit_feedback(
    State(pool): State<PgPool>,
    user: AuthUser,
    Path(review_id): Path<i32>,
    Json(req): Json<SubmitFeedbackRequest>,
) -> Result<Json<Review>> {
    let review = ReviewService::submit_feedback(&pool, review_id, &user.email, &req).await?;
    Ok(Json(review))
}
