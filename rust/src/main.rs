use review_api_rust::api::create_router;
use review_api_rust::infrastructure::{create_pool, Settings};
use tower_http::trace::TraceLayer;
use tracing_subscriber::{layer::SubscriberExt, util::SubscriberInitExt};

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    dotenvy::dotenv().ok();

    tracing_subscriber::registry()
        .with(
            tracing_subscriber::EnvFilter::try_from_default_env()
                .unwrap_or_else(|_| "review_api_rust=debug,tower_http=debug".into()),
        )
        .with(tracing_subscriber::fmt::layer())
        .init();

    let settings = Settings::from_env()
        .expect("Failed to load settings");

    let pool = create_pool(&settings.database_url).await?;

    tracing::info!("Connected to database");

    let app = create_router(pool)
        .layer(TraceLayer::new_for_http());

    let listener = tokio::net::TcpListener::bind(&settings.server_address())
        .await?;

    tracing::info!("Server listening on {}", listener.local_addr()?);

    axum::serve(listener, app).await?;

    Ok(())
}
