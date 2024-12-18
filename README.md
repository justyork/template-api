# API Template

A lightweight and scalable REST API server built with GoLang.

## Overview

This project is a template for building RESTful APIs using Go. It includes features such as JWT authentication, database
migrations, and Swagger documentation.

## Features

- **JWT Authentication**: Secure your endpoints with JSON Web Tokens.
- **Database Migrations**: Manage your database schema changes with ease.
- **Swagger Documentation**: Automatically generated API documentation.
- **Docker Support**: Easily containerize your application.
- **CI/CD Pipeline**: Automated build and deployment using GitHub Actions.

## Getting Started

### Prerequisites

- Go 1.23 or later
- Docker (optional, for containerization)
- SQLite (for local development)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/justyork/api-template.git
   cd api-template
   ```

2. Copy the example environment file:

   ```bash
   cp .env.example .env
   ```

3. Update `.env` with your configuration.

4. Run the application:

   ```bash
   go run cmd/main.go
   ```

### Running with Docker

1. Build the Docker image:

   ```bash
   docker build -t api-template .
   ```

2. Run the Docker container:

   ```bash
   docker-compose up
   ```

### Running Tests

Execute the following command to run tests:

```bash
go test ./... -v
```

## API Documentation

Access the Swagger UI at `http://localhost:8080/swagger/` to view the API documentation.

## Demo

Explore the live demo of the [deployed project](https://go-template-api.justyork.dev) and its [Swagger documentation](https://go-template-api.justyork.dev/swagger/index.html).

## CI/CD Pipeline

The project includes a GitHub Actions workflow for continuous integration and deployment. The pipeline is triggered on
pushes to the `main` branch and performs the following steps:

- Checkout code
- Setup Go environment
- Build the application
- Upload the binary to the server
- Run the application on the server

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE.md) file for details.

## Contact

For support, contact [yorkshp@gmail.com](mailto:yorkshp@gmail.com).