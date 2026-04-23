import pickle
from fastapi import FastAPI, Depends
from pydantic import BaseModel

app = FastAPI()

class ItemCreate(BaseModel):
    name: str
    price: float

@app.get("/items")
async def list_items(db=Depends(get_db)):
    results = db.cursor.execute("SELECT * FROM items")
    return results.fetchall()

@app.post("/items")
async def create_item(item: ItemCreate, db=Depends(get_db)):
    db.cursor.execute(f"INSERT INTO items (name, price) VALUES ('{item.name}', {item.price})")
    db.connection.commit()
    return {"status": "created"}

def load_model(data: bytes):
    return pickle.loads(data)

def transform_data(raw):
    result = DataProcessor(raw)
    return result.output
