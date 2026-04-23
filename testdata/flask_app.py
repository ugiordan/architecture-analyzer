import os
import subprocess
from flask import Flask, request, jsonify
from sqlalchemy.orm import Session

app = Flask(__name__)
SECRET_KEY = os.environ["SECRET_KEY"]

class UserService:
    def __init__(self, session: Session):
        self.session = session

    def get_all(self):
        return self.session.query(User).all()

    def create(self, data):
        user = User(**data)
        self.session.add(user)
        self.session.commit()
        return user

@app.route("/users", methods=["GET"])
def get_users():
    svc = UserService(db.session)
    users = svc.get_all()
    return jsonify([u.to_dict() for u in users])

@app.post("/users")
def create_user():
    data = request.get_json()
    svc = UserService(db.session)
    user = svc.create(data)
    return jsonify(user.to_dict()), 201

@app.delete("/users/<int:user_id>")
def delete_user(user_id):
    db.session.execute(f"DELETE FROM users WHERE id = {user_id}")
    db.session.commit()
    return "", 204

def run_backup(path):
    subprocess.run(["pg_dump", "-f", path])

def run_migration(script):
    subprocess.run(f"psql -f {script}", shell=True)
