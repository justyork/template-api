# Use lightweight Go image
FROM golang:1.23-alpine

WORKDIR /app

# Install dependencies
RUN apk add --no-cache sqlite

# Copy dependency files and download modules
COPY go.mod go.sum ./
RUN go mod download

# Copy application files
COPY . .

# Create .env from .env.example if .env does not exist
RUN if [ ! -f .env ]; then cp .env.example .env; fi

# Build the application
RUN go build -o api-template cmd/main.go

# Expose application port
EXPOSE 8080

# Run the application
CMD ["./api-template"]