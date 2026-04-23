use axum::{routing::get, routing::post, Router, Json, extract::State};
use sqlx::PgPool;
use serde::{Deserialize, Serialize};

#[derive(Serialize)]
struct User {
    id: i32,
    name: String,
}

#[derive(Deserialize)]
struct CreateUser {
    name: String,
}

async fn list_users(State(pool): State<PgPool>) -> Json<Vec<User>> {
    let users = sqlx::query_as!(User, "SELECT id, name FROM users")
        .fetch_all(&pool)
        .await
        .unwrap();
    Json(users)
}

async fn create_user(State(pool): State<PgPool>, Json(payload): Json<CreateUser>) -> Json<User> {
    let user = sqlx::query_as!(User, "INSERT INTO users (name) VALUES ($1) RETURNING id, name", payload.name)
        .fetch_one(&pool)
        .await
        .unwrap();
    Json(user)
}

unsafe fn raw_pointer_op(ptr: *const u8, len: usize) -> Vec<u8> {
    std::slice::from_raw_parts(ptr, len).to_vec()
}

extern "C" fn ffi_callback(data: *const u8) -> i32 {
    unsafe {
        if data.is_null() { return -1; }
        *data as i32
    }
}

fn build_router(pool: PgPool) -> Router {
    Router::new()
        .route("/users", get(list_users))
        .route("/users", post(create_user))
        .with_state(pool)
}
