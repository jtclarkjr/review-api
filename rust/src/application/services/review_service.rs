use sqlx::PgPool;

use crate::domain::models::{
    AssignReviewerRequest, CreateReviewRequest, Review, ReviewReviewer,
    SubmitFeedbackRequest, UpdateReviewRequest,
};
use crate::infrastructure::{AppError, EmployeeRepository, Result, ReviewRepository};

pub struct ReviewService;

impl ReviewService {
    pub async fn create(pool: &PgPool, req: &CreateReviewRequest) -> Result<Review> {
        ReviewRepository::create(pool, req).await
    }

    pub async fn get_all(pool: &PgPool) -> Result<Vec<Review>> {
        ReviewRepository::find_all(pool).await
    }

    pub async fn get_by_id(pool: &PgPool, id: i32) -> Result<Review> {
        ReviewRepository::find_by_id(pool, id).await
    }

    pub async fn get_assigned_reviews(pool: &PgPool, employee_email: &str) -> Result<Vec<Review>> {
        let employee = EmployeeRepository::find_by_email(pool, employee_email)
            .await?
            .ok_or_else(|| AppError::NotFound("Employee not found".to_string()))?;

        ReviewRepository::find_by_reviewer(pool, employee.id).await
    }

    pub async fn update(pool: &PgPool, id: i32, req: &UpdateReviewRequest) -> Result<Review> {
        ReviewRepository::update(pool, id, req).await
    }

    pub async fn assign_reviewer(
        pool: &PgPool,
        review_id: i32,
        req: &AssignReviewerRequest,
    ) -> Result<ReviewReviewer> {
        ReviewRepository::assign_reviewer(pool, review_id, req.reviewer_id).await
    }

    pub async fn submit_feedback(
        pool: &PgPool,
        review_id: i32,
        employee_email: &str,
        req: &SubmitFeedbackRequest,
    ) -> Result<Review> {
        let employee = EmployeeRepository::find_by_email(pool, employee_email)
            .await?
            .ok_or_else(|| AppError::NotFound("Employee not found".to_string()))?;

        let is_assigned = ReviewRepository::is_reviewer_assigned(pool, review_id, employee.id).await?;
        if !is_assigned {
            return Err(AppError::Forbidden);
        }

        ReviewRepository::add_comment(pool, review_id, &req.comment).await
    }
}
