import json
import pickle

def handle_request(request):
    body = request.data
    payload = json.loads(body)
    name = payload["name"]
    query = "SELECT * FROM users WHERE name = " + name
    return query

def simple_assignment():
    x = "hello"
    y = x
    return y

def field_access(user):
    email = user.email
    return email

class UserService:
    def create_user(self, data):
        name = data["name"]
        result = self.db.execute(name)
        return result
