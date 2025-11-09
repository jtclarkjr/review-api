use sqlx::PgPool;

use crate::domain::models::User;
use crate::infrastructure::error::Result;

pub struct UserRepository;

impl UserRepository {
    pub async fn find_by_email(pool: &PgPool, email: &str) -> Result<Option<User>> {
        let user = sqlx::query_as::<_, User>("SELECT * FROM users WHERE email = $1")
            .bind(email)
            .fetch_optional(pool)
            .await?;

        Ok(user)
    }
}
