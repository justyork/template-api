FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN apk add --no-cache sqlite

COPY . ./

RUN go build -o api-template cmd/main.go

COPY migrations /app/migrations

EXPOSE 8080

CMD ["./api-template"]