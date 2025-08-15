# Simple REST API

A basic RESTful API with CRUD operations for managing items.

## Features

- Create, Read, Update, and Delete items
- JSON data format
- RESTful endpoints

## Endpoints

- `GET /items` - Get all items
- `GET /items/{id}` - Get a specific item
- `POST /items` - Create a new item
- `PUT /items/{id}` - Update an existing item
- `DELETE /items/{id}` - Delete an item

## Setup

1. Navigate to the project directory:
   ```bash
   cd web-apis/simple-rest-api
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the server:
   ```bash
   go run .
   ```

## Usage

After starting the server, you can interact with the API using curl or any HTTP client:

```bash
# Get all items
curl http://localhost:8000/items

# Get a specific item
curl http://localhost:8000/items/1

# Create a new item
curl -X POST -H "Content-Type: application/json" -d '{"name":"Keyboard","price":50}' http://localhost:8000/items

# Update an item
curl -X PUT -H "Content-Type: application/json" -d '{"name":"Gaming Keyboard","price":100}' http://localhost:8000/items/1

# Delete an item
curl -X DELETE http://localhost:8000/items/1
```

## Data Structure

Items have the following structure:
```json
{
  "id": 1,
  "name": "Laptop",
  "price": 1000
}
