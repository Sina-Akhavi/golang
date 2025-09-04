# Use the official Go image as the base image for building
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application for Linux architecture
RUN GOOS=linux GOARCH=amd64 go build -o otp-service .

# Use a smaller base image for the final container
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Install glibc for dynamically linked Go binaries
RUN apk add --no-cache \
    libc6-compat \
    gcompat

# Copy the binary from the builder stage
COPY --from=builder /app/otp-service /app/otp-service

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./otp-service"]