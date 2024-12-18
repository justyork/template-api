# API Template

A lightweight and scalable REST API server built with GoLang and Docker.

## Features
- Simple and modular structure
- Dockerized for easy deployment
- Ready to scale

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

#### 3.	Access the API:
- http://localhost:8080

#### 4. Project Structure
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

#### 5. Technologies Used
- GoLang
- Docker
- Gorilla Mux