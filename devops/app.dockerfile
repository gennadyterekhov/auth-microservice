# Stage 1: Build the Go application using an Alpine-based Go image
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY . .
RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build  -o ./bin/server_linux_amd64  ./cmd/server


# Stage 2: Create a smaller image with the built binary
#FROM alpine:latest

# Set the working directory inside the container
#WORKDIR /app

## Copy the required files from the builder stage
#COPY --from=builder /app/bin/server_linux_amd64 .
#COPY --from=builder /app/.env .
#COPY --from=builder /app/go.mod .
#COPY --from=builder /app/certificates/server.crt ./certificates/server.crt
#COPY --from=builder /app/certificates/server.key ./certificates/server.key

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
#CMD ["/app/server_linux_amd64"]