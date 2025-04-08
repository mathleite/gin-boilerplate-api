# Gin Boilerplate API

`gin-boilerplate-api` is a Go-based microservice designed to handle client-related operations. This API provides endpoints for managing user data and authentication.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Examples](#examples)

## Features

- RESTful API for client management.
- Built with Go for high performance and scalability.
- Easy to integrate with other microservices.

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/mathleite/gin-boilerplate-api-api.git
   ```
2. Navigate to the project directory:
   ```bash
   cd gin-boilerplate-api-api
   ```
3. Create and fill .env:
   ```bash
   cp .env.example .env && \
   vim .env
   ```
4. Run the application:
   ```bash
   docker compose up --build --no-cache -d
   ```

## API Endpoints

### Base URL

```
http://localhost:8000/api/v1
```

### Endpoints

- **GET /health**: Application _health-check_.
- **POST /register**: Register a user.
- **POST /signin**: SignIn a user.
- **GET /transactions**: Access a authenticated endpoint _(has no logic here)_.

## Examples

### Create a User

**Request:**

```bash
curl --request POST \
  --url http://localhost:8000/api/v1/register \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "Boiler Plate",
	"email": "me.boilerplate@devlocal.com",
	"password": "123root"
}'
```

### SignIn

**Request:**

```bash
curl --request POST \
  --url http://localhost:8000/api/v1/auth/signin \
  --header 'Content-Type: application/json' \
  --data '{
	"email": "me.matheusleite@devlocal.com",
	"password": "123root"
}'
```

**Response:**

```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDEzMDUxMTcsImlhdCI6MTc0MTMwMzMxNywidXNlcl9pZCI6ImRiNGI5YWE4LTMwZDMtNDExMy05MjE2LTUwZTE0MTQ2OGQ1NSJ9.ZvUjK32lDg6OTFdNl-OTQDV_2ywcYLAWNWu3V23QYgo"
}
```

### "Fetch transactions"

**Request:**

```bash
curl --request GET \
  --url http://localhost:8000/api/v1/transactions \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDEzMDUxMTcsImlhdCI6MTc0MTMwMzMxNywidXNlcl9pZCI6ImRiNGI5YWE4LTMwZDMtNDExMy05MjE2LTUwZTE0MTQ2OGQ1NSJ9.ZvUjK32lDg6OTFdNl-OTQDV_2ywcYLAWNWu3V23QYgo'
```
