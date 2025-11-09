use axum::{
    routing::{delete, get, post, put},
    Router,
};
use sqlx::PgPool;

use crate::api::routes;

pub fn create_router(pool: PgPool) -> Router {
    Router::new()
        .route("/login", post(routes::login))
        .nest("/admin", admin_routes())
        .nest("/employee", employee_routes())
        .with_state(pool)
}

fn admin_routes() -> Router<PgPool> {
    Router::new()
        .route("/employees", post(routes::create_employee))
        .route("/employees", get(routes::get_employees))
        .route("/employees/:id", put(routes::update_employee))
        .route("/employees/:id", delete(routes::delete_employee))
        .route("/reviews", post(routes::create_review))
        .route("/reviews", get(routes::get_reviews))
        .route("/reviews/:id", put(routes::update_review))
        .route("/reviews/:review_id/assign", post(routes::assign_reviewer))
}

fn employee_routes() -> Router<PgPool> {
    Router::new()
        .route("/reviews", get(routes::get_assigned_reviews))
        .route("/reviews/:review_id/feedback", post(routes::submit_feedback))
}
