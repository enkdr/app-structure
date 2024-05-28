# Use the official Golang image as base
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY . .

# Set the environment variables for Go
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ENV DOCKER_ENV=true

# Build the Go application
RUN go build -o /app/bin/main ./cmd

# Expose port 3000 to the outside world
EXPOSE 8001

# Command to run the executable
CMD ["/app/bin/main"]

