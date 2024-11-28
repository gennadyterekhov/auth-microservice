FROM alpine

# Stage 1: Build the Go application using an Alpine-based Go image
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY . .
RUN go mod download

# install make
RUN apk add --no-cache make

# Run the build task using make
RUN make build


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