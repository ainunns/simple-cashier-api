# Simple Cashier API

A simple RESTful API for managing products, categories, and transactions in a cashier/POS system. Built with Go following clean architecture principles with proper separation of concerns using handlers, services, and repositories layers.

## Features

- **Product Management**: CRUD operations with search by name functionality
- **Category Management**: CRUD operations for product categories
- **Transaction Processing**: Checkout functionality with automatic stock management
- **Transaction Reports**: Daily and date-ranged transaction reports with best-selling products
- **PostgreSQL Database**: Persistent data storage with connection pooling
- **Clean Architecture**: Separated layers (handlers, services, repositories)
- **Environment Configuration**: Configurable via environment variables
- **JSON API**: RESTful endpoints with JSON responses

## Project Structure

```
cashier-api/
├── main.go                        # Application entry point and routing
├── go.mod                         # Go module dependencies
├── .env.example                   # Environment configuration
├── database/                      # Database connection
│   └── database.go                # PostgreSQL connection setup
├── models/                        # Data models
│   ├── product.go                 # Product models
│   ├── category.go                # Category model
│   └── transaction.go             # Transaction models
├── handlers/                      # HTTP handlers (presentation layer)
│   ├── product_handler.go         # Product HTTP handlers
│   ├── category_handler.go        # Category HTTP handlers
│   └── transaction_handler.go     # Transaction HTTP handlers
├── services/                      # Business logic layer
│   ├── product_service.go         # Product business logic
│   ├── category_service.go        # Category business logic
│   └── transaction_service.go     # Transaction business logic
└── repositories/                  # Data access layer
    ├── product_repository.go      # Product database operations
    ├── category_repository.go     # Category database operations
    └── transaction_repository.go  # Transaction database operations
```

## Prerequisites

- Go 1.23.2 or higher
- PostgreSQL database

## Installation

1. Clone the repository:

```bash
git clone <repository-url>
cd cashier-api
```

2. Install dependencies:

```bash
go mod download
```

3. Copy `.env.example` to `.env` file in the project root, then setup the application port and database connection:

```bash
cp .env.example .env
```

4. Build the application:

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

The server will start on the configured port (default: `http://0.0.0.0:8888`)

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

### Products

#### Get All Products

Get a list of all products with optional search by name.

**Endpoint:** `GET /api/products`

**Query Parameters:**

- `name` (optional): Filter products by name

**Example:** `GET /api/products?name=Indomie`

**Response:**

```json
[
  {
    "id": 1,
    "name": "Indomie Goreng",
    "price": 3500,
    "stock": 100,
    "category_id": 1,
    "category": {
      "id": 1,
      "name": "Makanan",
      "description": "Kategori untuk makanan"
    }
  }
]
```

#### Get Product by ID

Get details of a specific product including its category.

**Endpoint:** `GET /api/products/{id}`

**Example:** `GET /api/products/1`

**Response:**

```json
{
  "id": 1,
  "name": "Indomie Goreng",
  "price": 3500,
  "stock": 100,
  "category_id": 1,
  "category": {
    "id": 1,
    "name": "Makanan",
    "description": "Kategori untuk makanan"
  }
}
```

#### Create Product

Create a new product.

**Endpoint:** `POST /api/products`

**Request Body:**

```json
{
  "name": "Aqua 600ml",
  "price": 2000,
  "stock": 100,
  "category_id": 2
}
```

**Response:** `201 Created`

```json
{
  "id": 4,
  "name": "Aqua 600ml",
  "price": 2000,
  "stock": 100,
  "category_id": 2
}
```

#### Update Product

Update an existing product.

**Endpoint:** `PUT /api/products/{id}`

**Example:** `PUT /api/products/1`

**Request Body:**

```json
{
  "name": "Indomie Goreng Special",
  "price": 4000,
  "stock": 50,
  "category_id": 1
}
```

**Response:**

```json
{
  "id": 1,
  "name": "Indomie Goreng Special",
  "price": 4000,
  "stock": 50,
  "category_id": 1
}
```

#### Delete Product

Delete a product.

**Endpoint:** `DELETE /api/products/{id}`

**Example:** `DELETE /api/products/1`

**Response:**

```json
{
  "message": "Product deleted successfully"
}
```

---

### Categories

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
  "id": 3,
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
  "message": "Category deleted successfully"
}
```

---

### Transactions

#### Checkout

Process a transaction with multiple items. This endpoint automatically deducts stock and calculates totals.

**Endpoint:** `POST /api/checkout`

**Request Body:**

```json
{
  "items": [
    {
      "product_id": 1,
      "quantity": 2
    },
    {
      "product_id": 3,
      "quantity": 1
    }
  ]
}
```

**Response:**

```json
{
  "id": 1,
  "total_amount": 10000,
  "created_at": "2026-02-08T14:30:00Z",
  "details": [
    {
      "id": 1,
      "transaction_id": 1,
      "product_id": 1,
      "product_name": "Indomie Goreng",
      "quantity": 2,
      "subtotal": 7000
    },
    {
      "id": 2,
      "transaction_id": 1,
      "product_id": 3,
      "product_name": "Aqua 600ml",
      "quantity": 1,
      "subtotal": 3000
    }
  ]
}
```

#### Get Today's Transaction Report

Get transaction report for today including total revenue, transaction count, and best-selling products.

**Endpoint:** `GET /api/report/hari-ini`

**Response:**

```json
{
  "total_revenue": 150000,
  "total_transaksi": 12,
  "produk_terlaris": {
    "nama": "Indomie Goreng",
    "qty_terjual": 25
  }
}
```

#### Get Transaction Report by Date Range

Get transaction report for a specific date range.

**Endpoint:** `GET /api/report`

**Query Parameters:**

- `start_date` (optional): Start date in YYYY-MM-DD format
- `end_date` (optional): End date in YYYY-MM-DD format

**Example:** `GET /api/report?start_date=2026-02-01&end_date=2026-02-07`

**Response:**

```json
{
  "total_revenue": 500000,
  "total_transaksi": 45,
  "produk_terlaris": {
    "nama": "Indomie Goreng",
    "qty_terjual": 80
  }
}
```

## Testing with cURL

### Health Check

```bash
curl http://localhost:8888/health
```

### Products

```bash
# Get all products
curl http://localhost:8888/api/products

# Get all products with name filter
curl http://localhost:8888/api/products?name=Indomie

# Get product by ID
curl http://localhost:8888/api/products/1

# Create product
curl -X POST http://localhost:8888/api/products \
  -H "Content-Type: application/json" \
  -d '{"name":"Aqua 600ml","price":2000,"stock":100,"category_id":2}'

# Update product
curl -X PUT http://localhost:8888/api/products/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Indomie Goreng Special","price":4000,"stock":50,"category_id":1}'

# Delete product
curl -X DELETE http://localhost:8888/api/products/1
```

### Categories

```bash
# Get all categories
curl http://localhost:8888/api/categories

# Get category by ID
curl http://localhost:8888/api/categories/1

# Create category
curl -X POST http://localhost:8888/api/categories \
  -H "Content-Type: application/json" \
  -d '{"name":"Snack","description":"Kategori untuk snack"}'

# Update category
curl -X PUT http://localhost:8888/api/categories/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Makanan Berat","description":"Kategori untuk makanan berat"}'

# Delete category
curl -X DELETE http://localhost:8888/api/categories/1
```

### Transactions

```bash
# Checkout transaction
curl -X POST http://localhost:8888/api/checkout \
  -H "Content-Type: application/json" \
  -d '{
    "items": [
      {"product_id": 1, "quantity": 2},
      {"product_id": 3, "quantity": 1}
    ]
  }'

# Get today's transaction report
curl http://localhost:8888/api/report/hari-ini

# Get transaction report by date range
curl "http://localhost:8888/api/report?start_date=2026-02-01&end_date=2026-02-07"
```

## Technologies Used

- **Go 1.23.2** - Programming language
- **PostgreSQL** - Database (via Supabase)
- **lib/pq** - PostgreSQL driver for Go
- **Viper** - Configuration management
- **net/http** - HTTP server implementation
- **encoding/json** - JSON encoding and decoding

## Architecture

This project follows Clean Architecture principles with clear separation of concerns:

- **Handlers Layer**: HTTP request/response handling and routing
- **Services Layer**: Business logic and transaction management
- **Repositories Layer**: Data access and database operations
- **Models Layer**: Data structures and domain entities

## Data Models

### Product

```go
type Product struct {
    ID         int    `json:"id"`
    Name       string `json:"name"`
    Price      int    `json:"price"`
    Stock      int    `json:"stock"`
    CategoryID *int   `json:"category_id"`
}

type ProductDetail struct {
    ID         int       `json:"id"`
    Name       string    `json:"name"`
    Price      int       `json:"price"`
    Stock      int       `json:"stock"`
    CategoryID *int      `json:"category_id"`
    Category   *Category `json:"category,omitempty"`
}
```

### Category

```go
type Category struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}
```

### Transaction

```go
type Transaction struct {
    ID          int                 `json:"id"`
    TotalAmount int                 `json:"total_amount"`
    CreatedAt   time.Time           `json:"created_at"`
    Details     []TransactionDetail `json:"details"`
}

type TransactionDetail struct {
    ID            int    `json:"id"`
    TransactionID int    `json:"transaction_id"`
    ProductID     int    `json:"product_id"`
    ProductName   string `json:"product_name,omitempty"`
    Quantity      int    `json:"quantity"`
    Subtotal      int    `json:"subtotal"`
}
```

## Notes

- The API uses PostgreSQL for persistent data storage
- Connection pooling is configured with max 25 open connections and 5 idle connections
- All endpoints return JSON responses with appropriate HTTP status codes
- Stock is automatically managed during checkout transactions
- Transaction reports calculate revenue and identify best-selling products
