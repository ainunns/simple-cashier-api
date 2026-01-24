# Simple Cashier API

A simple RESTful API for managing products and categories in a cashier system. Built with Go following clean architecture principles with modular design.

## Features

- CRUD operations for Products (Produk)
- CRUD operations for Categories
- Thread-safe data storage using mutex
- Clean architecture with separate layers (models, handlers, storage)
- JSON API responses
- Health check endpoint

## Project Structure

```
simple-cashier-api/
├── main.go              # Application entry point and routing
├── go.mod               # Go module dependencies
├── models/              # Data models
│   ├── produk.go        # Product model definition
│   └── category.go      # Category model definition
├── handlers/            # HTTP handlers
│   ├── produk.go        # Product HTTP handlers
│   ├── category.go      # Category HTTP handlers
│   └── health.go        # Health check handler
└── storage/             # Data storage layer
    ├── produk.go        # Product storage with thread-safe operations
    └── category.go      # Category storage with thread-safe operations
```

## Prerequisites

- Go 1.23.2 or higher

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd simple-cashier-api
```

2. Build the application:
```bash
go build
```

## Running the Server

Start the server:
```bash
go run main.go
```

Or run the compiled binary:
```bash
./simple-cashier-api
```

The server will start on `http://localhost:8080`

## API Documentation

### Health Check

Check if the API is running.

**Endpoint:** `GET /health`

**Response:**
```json
{
  "status": "OK",
  "message": "API running"
}
```

---

### Products (Produk)

#### Get All Products
Get a list of all products.

**Endpoint:** `GET /api/produk`

**Response:**
```json
[
  {
    "id": 1,
    "nama": "Indomie Godog",
    "harga": 3500,
    "stok": 10
  },
  {
    "id": 2,
    "nama": "Vit 1000ml",
    "harga": 3000,
    "stok": 40
  }
]
```

#### Get Product by ID
Get details of a specific product.

**Endpoint:** `GET /api/produk/{id}`

**Example:** `GET /api/produk/1`

**Response:**
```json
{
  "id": 1,
  "nama": "Indomie Godog",
  "harga": 3500,
  "stok": 10
}
```

#### Create Product
Create a new product.

**Endpoint:** `POST /api/produk`

**Request Body:**
```json
{
  "nama": "Aqua 600ml",
  "harga": 2000,
  "stok": 100
}
```

**Response:** `201 Created`
```json
{
  "id": 4,
  "nama": "Aqua 600ml",
  "harga": 2000,
  "stok": 100
}
```

#### Update Product
Update an existing product.

**Endpoint:** `PUT /api/produk/{id}`

**Example:** `PUT /api/produk/1`

**Request Body:**
```json
{
  "nama": "Indomie Goreng",
  "harga": 4000,
  "stok": 20
}
```

**Response:**
```json
{
  "id": 1,
  "nama": "Indomie Goreng",
  "harga": 4000,
  "stok": 20
}
```

#### Delete Product
Delete a product.

**Endpoint:** `DELETE /api/produk/{id}`

**Example:** `DELETE /api/produk/1`

**Response:**
```json
{
  "message": "sukses delete"
}
```

---

### Categories (Kategori)

#### Get All Categories
Get a list of all categories.

**Endpoint:** `GET /api/categories`

**Response:**
```json
[
  {
    "id": 1,
    "name": "Makanan",
    "description": "Kategori untuk makanan"
  },
  {
    "id": 2,
    "name": "Minuman",
    "description": "Kategori untuk minuman"
  },
  {
    "id": 3,
    "name": "Sembako",
    "description": "Kategori untuk sembako"
  }
]
```

#### Get Category by ID
Get details of a specific category.

**Endpoint:** `GET /api/categories/{id}`

**Example:** `GET /api/categories/1`

**Response:**
```json
{
  "id": 1,
  "name": "Makanan",
  "description": "Kategori untuk makanan"
}
```

#### Create Category
Create a new category.

**Endpoint:** `POST /api/categories`

**Request Body:**
```json
{
  "name": "Snack",
  "description": "Kategori untuk snack"
}
```

**Response:** `201 Created`
```json
{
  "id": 4,
  "name": "Snack",
  "description": "Kategori untuk snack"
}
```

#### Update Category
Update an existing category.

**Endpoint:** `PUT /api/categories/{id}`

**Example:** `PUT /api/categories/1`

**Request Body:**
```json
{
  "name": "Makanan Berat",
  "description": "Kategori untuk makanan berat"
}
```

**Response:**
```json
{
  "id": 1,
  "name": "Makanan Berat",
  "description": "Kategori untuk makanan berat"
}
```

#### Delete Category
Delete a category.

**Endpoint:** `DELETE /api/categories/{id}`

**Example:** `DELETE /api/categories/1`

**Response:**
```json
{
  "message": "sukses delete"
}
```

## Testing with cURL

### Products

```bash
# Get all products
curl http://localhost:8080/api/produk

# Get product by ID
curl http://localhost:8080/api/produk/1

# Create product
curl -X POST http://localhost:8080/api/produk \
  -H "Content-Type: application/json" \
  -d '{"nama":"Aqua 600ml","harga":2000,"stok":100}'

# Update product
curl -X PUT http://localhost:8080/api/produk/1 \
  -H "Content-Type: application/json" \
  -d '{"nama":"Indomie Goreng","harga":4000,"stok":20}'

# Delete product
curl -X DELETE http://localhost:8080/api/produk/1
```

### Categories

```bash
# Get all categories
curl http://localhost:8080/api/categories

# Get category by ID
curl http://localhost:8080/api/categories/1

# Create category
curl -X POST http://localhost:8080/api/categories \
  -H "Content-Type: application/json" \
  -d '{"name":"Snack","description":"Kategori untuk snack"}'

# Update category
curl -X PUT http://localhost:8080/api/categories/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Makanan Berat","description":"Kategori untuk makanan berat"}'

# Delete category
curl -X DELETE http://localhost:8080/api/categories/1
```

### Health Check

```bash
curl http://localhost:8080/health
```

## Technologies Used

- **Go 1.23.2** - Programming language
- **net/http** - HTTP server and client implementation
- **encoding/json** - JSON encoding and decoding

## Data Models

### Product (Produk)
```go
type Produk struct {
    ID    int    `json:"id"`
    Nama  string `json:"nama"`
    Harga int    `json:"harga"`
    Stok  int    `json:"stok"`
}
```

### Category (Kategori)
```go
type Category struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}
```

## Notes

- This API uses in-memory storage. Data will be lost when the server restarts.
- The implementation uses `sync.RWMutex` for thread-safe concurrent access to data.
- IDs are automatically generated based on the current number of items.
- All endpoints return JSON responses with appropriate HTTP status codes.
