# Huma Project Starter

A Go-based API service starter kit built with the Huma framework, demonstrating best practices for building and maintaining scalable REST APIs.

## 📋 Project Overview

This project provides a foundation for building RESTful APIs with Go. It includes:

- A clean architecture design
- API endpoint examples with request/response validation
- Database integration using Bun ORM
- Error handling and logging with Sentry integration
- Environment-based configuration

## 🏗️ Project Structure

```
apps/
  api-service/           # Main API service
    app/                 # Application core
      greeting/          # Greeting module (example)
      internal-lib/      # Internal libraries
        snowflake/       # Snowflake ID implementation
        utils/           # Utility functions
      setup/             # App configuration and setup
      shared/            # Shared resources
      user/              # User module
    migrations/          # Database migrations
    routes/              # API route definitions
    Dockerfile           # Container definition
    example.env          # Environment variable example
    go.mod               # Go module definition
    main.go              # Application entry point
```

## 🚀 Features

- RESTful API endpoints using Huma framework
- Chi router for HTTP request handling
- Database integration using Bun ORM
- Snowflake ID generation for distributed systems
- Environment-based configuration
- Dependency injection with Wire
- Error tracking with Sentry

## 🛠️ API Endpoints

The service currently provides the following endpoints:

- `GET /greeting/{name}`: Returns a greeting message for the provided name
- `POST /reviews`: Endpoint for submitting reviews (rated 1-5)

## 🔧 Getting Started

### Prerequisites

- Go 1.16+
- PostgreSQL (or compatible database)

### Environment Setup

Copy the example environment file and adjust as needed:

```bash
cp example.env .env
```

Required environment variables:
- `ENVIRONMENT`: Set to `development`, `staging`, or `production`
- `PORT`: HTTP port to listen on
- `GRPC_PORT`: gRPC port (if enabled)
- `API_KEY`: API key for authentication
- `DEBUG`: Set to `true` for verbose logging

### Running Locally

```bash
go run main.go
```

### Building for Production

```bash
go build -o api-service
```

## 🐳 Docker Support

This project includes Docker support for containerized deployment.

Build the container:
```bash
docker build -t huma-api-service .
```

Run the container:
```bash
docker run -p 8080:8080 --env-file .env huma-api-service
```

## 📚 Documentation

API documentation is automatically generated through the Huma framework and available at the `/docs` endpoint when the server is running.

## 🧪 Testing

Run the tests with:

```bash
go test ./...
```

## 📦 Dependencies

- [Huma](https://github.com/danielgtaylor/huma/): API framework
- [Chi Router](https://github.com/go-chi/chi): HTTP routing
- [Bun](https://github.com/uptrace/bun): SQL ORM
- [Sentry](https://github.com/getsentry/sentry-go): Error tracking
- [Wire](https://github.com/google/wire): Dependency injection

## 📝 License

[MIT License]
