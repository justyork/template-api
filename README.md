# API Template

A lightweight and scalable REST API server built with GoLang and Docker.

## Features
- Simple and modular structure
- Dockerized for easy deployment
- Ready to scale

### Authentication

This project uses JSON Web Tokens (JWT) for authentication.

#### Endpoints
- `POST /login`: Accepts JSON payload with `username` and `password`, returns a JWT token in a cookie.
- `GET /protected`: A protected route that requires a valid JWT token.

#### Example Usage
Login:
```bash
curl -X POST http://localhost:8080/login -d '{"username":"testuser", "password":"password"}' -H "Content-Type: application/json"
```
Access Protected Route:
```bash
curl -X GET http://localhost:8080/protected --cookie "token=your_token"
```

### Environment Configuration

This project uses environment variables for configuration.

#### `.env.example`
The file `.env.example` is provided as a template. It contains default values for required variables:

```env
DATABASE_URL=sqlite://db.sqlite
JWT_SECRET=your_secret_key
APP_PORT=8080
```

How it Works
- When building the Docker image, if .env is not present, it will be automatically created from .env.example.
- You can customize .env after the build if needed.

Manually Creating .env

If you’re running the application locally, create `.env` by copying `.env.example`:
```bash
cp .env.example .env
```

Then, edit the file with your desired configuration.
---

## Getting Started

### Prerequisites
- Docker installed on your system

### Running the Project
#### 1. Clone the repository:
```bash
    git clone https://github.com/justyork/api-template.git
    cd api-template
```
#### 2.	Build and run the application:
```bash
    docker-compose up --build 
```
#### Running Migrations
Migrations are automatically applied when you start the application. Ensure the `migrations/` folder is included in your setup.

To run migrations manually, use:
```bash
migrate -path ./migrations -database sqlite3://db.sqlite up
```

To rollback:
```bash
migrate -path ./migrations -database sqlite3://db.sqlite down
```
   
#### 4.	Access the API:
- http://localhost:8080

#### 5. Project Structure
```
api-template/
├── cmd/            # Entry point
│   └── main.go
├── internal/       # Application logic
│   ├── handlers/   # HTTP-handles
│   ├── services/   # Data logic
│   └── models/     # Data models
├── pkg/            # Common libriaries
├── configs/        # Configuration fields
├── migrations/     # SQL-migrations
├── Dockerfile      
├── docker-compose.yml 
├── go.mod          # Go dependences
└── README.md       # Project description
```

#### 6. Technologies Used
- GoLang
- Docker
- Gorilla Mux