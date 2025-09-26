# E-commerce API

This is a simple e-commerce API written in Go using the Gin framework.

## About the Project

This project provides a RESTful API for a basic e-commerce platform. It includes features like user authentication, product management, and more.

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

- Go (version 1.23.0 or later)
- PostgreSQL

### Installation

1.  Clone the repo
    ```sh
    git clone https://github.com/your_username/e-commerce.git
    ```
2.  Install Go packages
    ```sh
    go mod download
    ```
3.  Set up your environment variables by creating a `.env` file in the root directory.
    ```
    JWT_SECRET=your_jwt_secret
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=user
    DB_PASSWORD=mysecretpassword
    DB_NAME=ecommerce
    PORT=8080
    ```
4.  Run the application
    ```sh
    go run cmd/main.go
    ```

## API Endpoints

### Authentication

-   `POST /auth/register`: Register a new user.
-   `POST /auth/login`: Login a user and get a JWT token.

### Users

-   `GET /user/me`: Get the profile of the currently logged-in user (requires authentication).
-   `PUT /user/me`: Update the profile of the currently logged-in user (requires authentication).

### Products

-   `GET /products`: Get a list of all products.
-   `GET /products/:id`: Get a single product by its ID.
-   `POST /products`: Create a new product (requires admin privileges).
-   `PUT /products/:id`: Update a product (requires admin privileges).
-   `DELETE /products/:id`: Delete a product (requires admin privileges).

## Database Models

The application uses the following database models:

### User

| Field     | Type      | Description                |
|-----------|-----------|----------------------------|
| ID        | `uint`    | Primary Key                |
| Name      | `string`  | User's name                |
| Email     | `string`  | User's email (unique)      |
| Password  | `string`  | User's password (hashed)   |
| Role      | `string`  | User's role (e.g., `user`, `admin`) |
| CreatedAt | `time.Time` | Time of user creation      |

### Product

| Field       | Type      | Description              |
|-------------|-----------|--------------------------|
| ID          | `uint`    | Primary Key              |
| Name        | `string`  | Product's name           |
| Description | `string`  | Product's description    |
| Price       | `uint`    | Product's price          |
| Stock       | `uint`    | Product's stock quantity |
| Category    | `string`  | Product's category       |
| CreatedAt   | `time.Time` | Time of product creation |

### Order

| Field       | Type      | Description              |
|-------------|-----------|--------------------------|
| ID          | `uint`    | Primary Key              |
| UserId      | `uint`    | Foreign key to User      |
| TotalAmount | `uint`    | Total amount of the order|
| Status      | `string`  | Order status (e.g., `pending`, `paid`, `shipped`) |
| PaymentId   | `string`  | Payment ID from the payment provider |
| CreatedAt   | `time.Time` | Time of order creation   |

### OrderItem

| Field     | Type   | Description           |
|-----------|--------|-----------------------|
| ID        | `uint` | Primary Key           |
| OrderId   | `uint` | Foreign key to Order  |
| ProductId | `uint` | Foreign key to Product|
| Quantity  | `uint` | Quantity of the product |
| Price     | `uint` | Price of the product at the time of purchase |

### Cart

| Field     | Type   | Description           |
|-----------|--------|-----------------------|
| ID        | `uint` | Primary Key           |
| UserId    | `uint` | Foreign key to User   |
| ProductID | `uint` | Foreign key to Product|
| Quantity  | `uint` | Quantity of the product in the cart |
