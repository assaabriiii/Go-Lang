# Go Web Application with User and Product Management

This is a web application built with Go and Gin framework that provides REST API endpoints for managing users and products.

## Features

- User Management
  - List all users
  - Add new users
  - Get user by ID
  - View users in HTML format
- Product Management
  - List all products
  - Add new products
  - Get product by ID

## API Endpoints

### Users

```http
GET /users # Get all users
POST /users # Create a new user
GET /users/:id # Get user by ID
GET /users/page # View users page (HTML)
```

### Products

```http
GET /products # Get all products
POST /products # Create a new product
GET /products/:id # Get product by ID
```

## Data Models

### User

```go
type User struct {
    ID int json:"id"
    Name string json:"name"
    Email string json:"email"
    Products []Product json:"products"
}
```

### Product

```go
type Product struct {
ID int json:"id"
Name string json:"name"
Price float64 json:"price"
UserID int json:"user_id,omitempty"
}
```


## Sample Data

The application comes pre-loaded with sample data including:
- 2 users (Alice and Bob)
- 4 products (Laptop, Phone, Tablet, and Headphones)

## Error Handling

The API returns appropriate HTTP status codes:
- 200: Successful operations
- 400: Bad request (invalid input)
- 404: Resource not found

## Getting Started

1. Make sure you have Go installed on your system
2. Clone this repository
3. Install dependencies:
   ```bash
   go mod download
   ```
4. Run the application:
   ```bash
   go run main.go
   ```

## Dependencies

- [Gin Web Framework](https://github.com/gin-gonic/gin) - Web framework for Go