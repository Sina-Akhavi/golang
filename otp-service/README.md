# OTP Service

## Overview

The OTP Service is a backend application built with Go and Gin. It provides:
- OTP generation and validation for authentication.
- JWT-based token authentication.
- Redis for temporary OTP storage.

---

## Features

- **Public Endpoints**:
  - `/request-otp`: Generate OTP for a phone number.
  - `/validate-otp`: Validate OTP and receive JWT token.
- **Protected Endpoints (JWT Required)**:
  - `/user`: Get user details.
  - `/users`: List users (paginated).
  - `/users` (POST): Create user.
- **Swagger Documentation**: `/swagger/`

---

## Prerequisites

- [Go (1.20+)](https://golang.org/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## How to Run Locally

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/<your-username>/otp-service.git
   cd otp-service
   ```

2. **Copy Environment File**:
   ```bash
   cp .env.example .env
   ```
   Edit `.env` as needed.

3. **Start Redis** (if not using Docker Compose):
   ```bash
   docker run -d --name redis -p 6379:6379 redis
   ```

4. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

5. **Run the Application**:
   ```bash
   go run main.go
   ```

6. **Access Swagger UI**:
   [http://localhost:8080/swagger/](http://localhost:8080/swagger/)

---

## How to Run with Docker

1. **Build and Start Services**:
   ```bash
   docker-compose up --build
   ```

2. **Access Swagger UI**:
   [http://localhost:8080/swagger/](http://localhost:8080/swagger/)

---

## Example API Requests & Responses

### 1. Request OTP

**Request**
```bash
curl -X POST http://localhost:8080/request-otp \
  -H "Content-Type: application/json" \
  -d '{"phone": "+1234567890"}'
```

**Response**
```json
{
  "message": "OTP sent successfully"
}
```

---

### 2. Validate OTP

**Request**
```bash
curl -X POST http://localhost:8080/validate-otp \
  -H "Content-Type: application/json" \
  -d '{"phone": "+1234567890", "otp": "123456"}'
```

**Response**
```json
{
  "token": "<jwt_token>"
}
```

---

### 3. Get User (JWT Required)

**Request**
```bash
curl -X GET http://localhost:8080/user \
  -H "Authorization: Bearer <jwt_token>"
```

**Response**
```json
{
  "id": 1,
  "phone": "+1234567890",
  "name": "John Doe"
}
```

---

## API Endpoints

| Method | Endpoint         | Description                      | Auth Required |
|--------|------------------|----------------------------------|--------------|
| POST   | `/request-otp`   | Request OTP for phone number     | No           |
| POST   | `/validate-otp`  | Validate OTP and get JWT token   | No           |
| GET    | `/user`          | Get user details                 | Yes          |
| GET    | `/users`         | List users (paginated)           | Yes          |
| POST   | `/users`         | Create new user                  | Yes          |

---

## Environment Variables

| Variable         | Description                |
|------------------|---------------------------|
| `REDIS_URL`      | Redis connection string   |
| `JWT_SECRET`     | Secret for JWT signing    |
| `PORT`           | Server port (default 8080)|

---

## Testing

To run unit tests:
```bash
go test ./...
```

---

## License

MIT License.

---

## Contributing

Pull requests are welcome! For major changes, open an issue first.

---

## Contact

For questions or support, please open an