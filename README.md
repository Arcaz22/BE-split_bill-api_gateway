<h1 align="center">
  <a href="https://go.dev/" target="blank"><img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original-wordmark.svg" height="100" alt="Golang" /></a>
  <a href="https://gin-gonic.com/" target="blank"><img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" height="100" alt="Gin Framework" /></a>
</h1>

<p align="center">API Gateway using Golang and Gin Framework</p>

## Description

This project is an API Gateway built with Go and Gin Framework, featuring:

- 🚀 High-performance proxy routing
- 🔐 Authentication middleware
- 🛡️ Circuit breaker pattern
- 📊 Rate limiting
- 📝 Request logging
- 🌐 CORS support
- 🔄 Load balancing ready
- 📡 Service discovery ready
- 🎯 Health checking
- ⚡ Fast response time

## Prerequisites

- Go >= 1.21
- Auth Service running on port 3000

## Getting Started

```bash
# Clone repository
git clone https://github.com/yourusername/api-gateway.git

# Install dependencies
cd api-gateway
go mod tidy

# Run the gateway
go run cmd/main.go
```

## Features

### Middleware
- Authentication checking
- Rate limiting (100 req/sec)
- Circuit breaker for service protection
- CORS handling
- Request logging

### Routing
- Auth service forwarding
- User management
- Role management
- Health check endpoint

### Monitoring
- Request logging
- Latency tracking
- Status code monitoring
- Circuit breaker state

## Configuration

Example config.yaml:
```yaml
server:
  port: "8080"
services:
  auth:
    url: "http://localhost:3000/v1"
```

## Available Routes

### Auth Routes
- POST /auth/signup
- POST /auth/signin
- POST /auth/refresh-token
- POST /auth/logout
- POST /auth/google
- GET /auth/google/callback

### User Routes
- GET /user/current-user
- GET /user/get-all-user
- POST /user/add-profile
- POST /user/add-avatar-profile
- PUT /user/update-profile
- PUT /user/update-avatar-profile

### Role Routes
- GET /role/get-all-role
- POST /role/create-role
- POST /role/add-user-role
