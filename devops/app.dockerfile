# Stage 1: Build the Go application using an Alpine-based Go image
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY . .
RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build  -o ./cmd/server/server_linux_amd64  ./cmd/server


# Stage 2: Create a smaller image with the built binary
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/cmd/server/server_linux_amd64 .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["/app/server_linux_amd64"]