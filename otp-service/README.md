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