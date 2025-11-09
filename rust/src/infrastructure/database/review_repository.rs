use sqlx::PgPool;

use crate::domain::models::{CreateReviewRequest, Review, ReviewReviewer, UpdateReviewRequest};
use crate::infrastructure::error::{AppError, Result};

pub struct ReviewRepository;

impl ReviewRepository {
    pub async fn create(pool: &PgPool, req: &CreateReviewRequest) -> Result<Review> {
        let review = sqlx::query_as::<_, Review>(
            "INSERT INTO reviews (employee_id, performance_review) VALUES ($1, $2) RETURNING *",
        )
        .bind(req.employee_id)
        .bind(&req.performance_review)
        .fetch_one(pool)
        .await?;

        Ok(review)
    }

    pub async fn find_all(pool: &PgPool) -> Result<Vec<Review>> {
        let reviews = sqlx::query_as::<_, Review>("SELECT * FROM reviews")
            .fetch_all(pool)
            .await?;

        Ok(reviews)
    }

    pub async fn find_by_id(pool: &PgPool, id: i32) -> Result<Review> {
        let review = sqlx::query_as::<_, Review>("SELECT * FROM reviews WHERE id = $1")
            .bind(id)
            .fetch_one(pool)
            .await?;

        Ok(review)
    }

    pub async fn find_by_reviewer(pool: &PgPool, reviewer_id: i32) -> Result<Vec<Review>> {
        let reviews = sqlx::query_as::<_, Review>(
            r#"
            SELECT r.* FROM reviews r
            INNER JOIN review_reviewers rr ON r.id = rr.review_id
            WHERE rr.reviewer_id = $1
            "#,
        )
        .bind(reviewer_id)
        .fetch_all(pool)
        .await?;

        Ok(reviews)
    }

    pub async fn update(pool: &PgPool, id: i32, req: &UpdateReviewRequest) -> Result<Review> {
        if let Some(ref performance_review) = req.performance_review {
            let review = sqlx::query_as::<_, Review>(
                "UPDATE reviews SET performance_review = $1 WHERE id = $2 RETURNING *",
            )
            .bind(performance_review)
            .bind(id)
            .fetch_one(pool)
            .await?;

            Ok(review)
        } else {
            Err(AppError::BadRequest("No fields to update".to_string()))
        }
    }

    pub async fn add_comment(pool: &PgPool, review_id: i32, comment: &str) -> Result<Review> {
        let review = sqlx::query_as::<_, Review>(
            "UPDATE reviews SET comments = array_append(comments, $1) WHERE id = $2 RETURNING *",
        )
        .bind(comment)
        .bind(review_id)
        .fetch_one(pool)
        .await?;

        Ok(review)
    }

    pub async fn assign_reviewer(
        pool: &PgPool,
        review_id: i32,
        reviewer_id: i32,
    ) -> Result<ReviewReviewer> {
        let reviewer = sqlx::query_as::<_, ReviewReviewer>(
            "INSERT INTO review_reviewers (review_id, reviewer_id) VALUES ($1, $2) RETURNING *",
        )
        .bind(review_id)
        .bind(reviewer_id)
        .fetch_one(pool)
        .await?;

        Ok(reviewer)
    }

    pub async fn is_reviewer_assigned(
        pool: &PgPool,
        review_id: i32,
        reviewer_id: i32,
    ) -> Result<bool> {
        let result = sqlx::query_as::<_, ReviewReviewer>(
            "SELECT * FROM review_reviewers WHERE review_id = $1 AND reviewer_id = $2",
        )
        .bind(review_id)
        .bind(reviewer_id)
        .fetch_optional(pool)
        .await?;

        Ok(result.is_some())
    }
}
