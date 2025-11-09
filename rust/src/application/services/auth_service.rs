use sqlx::PgPool;

use crate::domain::models::{LoginRequest, LoginResponse};
use crate::infrastructure::{create_jwt, verify_password, AppError, Result, UserRepository};

pub struct AuthService;

impl AuthService {
    pub async fn login(pool: &PgPool, req: &LoginRequest) -> Result<LoginResponse> {
        let user = UserRepository::find_by_email(pool, &req.email)
            .await?
            .ok_or(AppError::Unauthorized)?;

        let valid = verify_password(&req.password, &user.password)?;
        if !valid {
            return Err(AppError::Unauthorized);
        }

        let token = create_jwt(&user.email, &user.role)?;

        Ok(LoginResponse { token })
    }
}
