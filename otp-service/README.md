# OTP Service

## Overview

The OTP Service is a simple backend application built with Go and the Gin web framework. It provides the following features:
- Generate and validate OTP (One-Time Password) for user authentication.
- Token-based authentication using JWT (JSON Web Token).
- Redis is used to store temporary OTPs for validation.

The service is designed to support authentication for mobile apps, web apps, or other clients. Protected routes require a valid JWT token for access.

---

## Features

1. **Public Endpoints**:
   - `/request-otp`: Generate an OTP for a given phone number.
   - `/validate-otp`: Validate the OTP and receive a JWT token for further authentication.

2. **Protected Endpoints (JWT Required)**:
   - `/user`: Get details of a single user.
   - `/users`: Retrieve a paginated list of users.
   - `/users` (POST): Create a new user.

3. **Swagger Documentation**:
   - Explore and test the API using the Swagger UI at `/swagger/`.

---

## Prerequisites

Before running the project, ensure you have the following installed:
- [Go (1.20+)](https://golang.org/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## Running the Project

### **Option 1: Run Locally**

To run the project locally, follow these steps:

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/<your-username>/otp-service.git
   cd otp-service
   ```

2. **Copy the Example Environment File**:
   ```bash
   cp .env.example .env
   ```
   Edit `.env` to set your environment variables (e.g., Redis connection, JWT secret).

3. **Start Redis (if not using Docker Compose)**:
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
   Open [http://localhost:8080/swagger/](http://localhost:8080/swagger/) in your browser.

---

### **Option 2: Run with Docker Compose**

1. **Start Services**:
   ```bash
   docker-compose up --build
   ```

2. **Access Swagger UI**:
   Open [http://localhost:8080/swagger/](http://localhost:8080/swagger/) in your browser.

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

This project is licensed under the MIT License.

---

## Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

---

## Contact

For questions or support, please open an issue or contact the