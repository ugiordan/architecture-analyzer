use actix_web::{get, post, web, HttpResponse, Responder};
use diesel::prelude::*;
use std::env;

#[get("/health")]
async fn health_check() -> impl Responder {
    HttpResponse::Ok().body("healthy")
}

#[post("/items")]
async fn create_item(pool: web::Data<DbPool>, item: web::Json<NewItem>) -> impl Responder {
    let conn = pool.get().expect("Failed to get DB connection");
    diesel::insert_into(items::table)
        .values(&*item)
        .execute(&conn)
        .expect("Failed to insert item");
    HttpResponse::Created().finish()
}

#[get("/items")]
async fn list_items(pool: web::Data<DbPool>) -> impl Responder {
    let conn = pool.get().expect("Failed to get DB connection");
    let results = items::table.load::<Item>(&conn).expect("Failed to load items");
    HttpResponse::Ok().json(results)
}

fn get_api_key() -> String {
    env::var("API_KEY").expect("API_KEY must be set")
}

fn get_db_url() -> String {
    std::env::var("DATABASE_URL").expect("DATABASE_URL must be set")
}
