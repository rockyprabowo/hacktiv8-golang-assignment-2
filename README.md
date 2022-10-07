# Hacktiv8 Scalable Web Service with Golang - Assignment 2

This is the second assignment for Scalable Web Service with Golang course by Hacktiv8.

This project is a REST API back end application made with Gin HTTP web framework, GORM for object–relational mapping backed by a PostgreSQL database, and swaggo for the API endpoints documentation.

## Usage

* Copy `.env.example` to `.env` and edit `.env` with your database configurations and listen/serve address.
* Run with `go run .` or build the application first with `go build .` if you wish.
* You're ready to go!

## Endpoints

* `GET    /orders` Fetch all orders
* `POST   /orders` Create order
* `GET    /orders/:id` Fetch an order
* `PUT    /orders/:id` Update an order
* `DELETE /orders/:id` Delete an order
* `POST   /orders/__prune__` Undocumented: Prune the database tables
* `GET    /swagger/*any` Swagger/OpenAPI 2.0 documentations
* `GET    /swagger/index.html` Swagger/OpenAPI 2.0 documentations index page


## Changelog

### 1.1
* Fix Swagger/OpenAPI 2.0 documentation page issues
* Internal package name change
* Update README.md

### 1.0
* Initial release
