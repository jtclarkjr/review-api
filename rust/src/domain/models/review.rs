use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use sqlx::FromRow;
use utoipa::ToSchema;

#[derive(Debug, FromRow, Serialize, Deserialize, ToSchema)]
pub struct Review {
    pub id: i32,
    pub employee_id: i32,
    pub performance_review: String,
    pub comments: Vec<String>,
    pub created_at: DateTime<Utc>,
}

#[derive(Debug, FromRow, Serialize, Deserialize, ToSchema)]
pub struct ReviewReviewer {
    pub id: i32,
    pub review_id: i32,
    pub reviewer_id: i32,
}

#[derive(Debug, FromRow, Serialize, Deserialize, ToSchema)]
pub struct Feedback {
    pub id: i32,
    pub review_id: i32,
    pub submitted: bool,
}

#[derive(Debug, Deserialize, ToSchema)]
pub struct CreateReviewRequest {
    pub employee_id: i32,
    pub performance_review: String,
}

#[derive(Debug, Deserialize, ToSchema)]
pub struct UpdateReviewRequest {
    pub performance_review: Option<String>,
}

#[derive(Debug, Deserialize, ToSchema)]
pub struct AssignReviewerRequest {
    pub reviewer_id: i32,
}

#[derive(Debug, Deserialize, ToSchema)]
pub struct SubmitFeedbackRequest {
    pub comment: String,
}
