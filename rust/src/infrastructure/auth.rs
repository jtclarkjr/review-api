use axum::{
    async_trait,
    extract::FromRequestParts,
    http::request::Parts,
    RequestPartsExt,
};
use jsonwebtoken::{decode, encode, DecodingKey, EncodingKey, Header, Validation};
use serde::{Deserialize, Serialize};
use std::env;

use crate::domain::models::Role;
use crate::infrastructure::error::{AppError, Result};

#[derive(Debug, Serialize, Deserialize)]
pub struct Claims {
    pub sub: String,
    pub role: Role,
    pub exp: usize,
}

pub struct AuthUser {
    pub email: String,
    pub role: Role,
}

#[async_trait]
impl<S> FromRequestParts<S> for AuthUser
where
    S: Send + Sync,
{
    type Rejection = AppError;

    async fn from_request_parts(parts: &mut Parts, _state: &S) -> Result<Self> {
        let bearer = parts
            .extract::<axum_extra::TypedHeader<axum_extra::headers::Authorization<axum_extra::headers::authorization::Bearer>>>()
            .await
            .map_err(|_| AppError::Unauthorized)?;

        let token = bearer.token();
        let claims = decode_jwt(token)?;

        Ok(AuthUser {
            email: claims.sub,
            role: claims.role,
        })
    }
}

pub fn get_jwt_secret() -> String {
    env::var("JWT_SECRET").unwrap_or_else(|_| "secret".to_string())
}

pub fn create_jwt(email: &str, role: &Role) -> Result<String> {
    let expiration = chrono::Utc::now()
        .checked_add_signed(chrono::Duration::hours(24))
        .expect("valid timestamp")
        .timestamp() as usize;

    let claims = Claims {
        sub: email.to_string(),
        role: role.clone(),
        exp: expiration,
    };

    let secret = get_jwt_secret();
    let token = encode(
        &Header::default(),
        &claims,
        &EncodingKey::from_secret(secret.as_bytes()),
    )?;

    Ok(token)
}

pub fn decode_jwt(token: &str) -> Result<Claims> {
    let secret = get_jwt_secret();
    let token_data = decode::<Claims>(
        token,
        &DecodingKey::from_secret(secret.as_bytes()),
        &Validation::default(),
    )?;

    Ok(token_data.claims)
}

pub fn verify_password(password: &str, hash: &str) -> Result<bool> {
    let valid = bcrypt::verify(password, hash)?;
    Ok(valid)
}

pub fn require_admin(user: &AuthUser) -> Result<()> {
    if user.role != Role::Admin {
        return Err(AppError::Forbidden);
    }
    Ok(())
}
