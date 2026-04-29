use std::io::Read;

fn handle_request(request: &str) -> String {
    let body = request.to_string();
    let payload: serde_json::Value = serde_json::from_str(&body).unwrap();
    let name = payload["name"].as_str().unwrap();
    let query = format!("SELECT * FROM users WHERE name = {}", name);
    query
}

fn simple_assignment() -> String {
    let x = String::from("hello");
    let y = x;
    return y;
}

fn field_access(user: &User) -> String {
    let email = user.email.clone();
    email
}

struct User {
    email: String,
}
