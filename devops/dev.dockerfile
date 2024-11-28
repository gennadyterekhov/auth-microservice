# Stage 1: Build the Go application using an Alpine-based Go image
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY . .
RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build  -o ./cmd/server/server_linux_amd64  ./cmd/server

EXPOSE 8080

# Command to run the application
CMD ["/app/cmd/server/server_linux_amd64"]
