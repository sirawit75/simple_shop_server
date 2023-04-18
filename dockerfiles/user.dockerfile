# Use the official Golang image as the builder image
FROM golang:1.19-alpine as builder

# Set the working directory to the project directory
WORKDIR /app

# Copy the Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application code into the container
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/user/main.go

# Use the official Alpine image as the production image
FROM alpine:latest

# Set the working directory to the project directory
WORKDIR /app

# Copy the binary from the builder image
COPY --from=builder /app/main ./
COPY ./cmd/user/user.env .

# Expose the port used by the application
EXPOSE 8080

# Set the default command to run the binary
CMD ["./main"]
