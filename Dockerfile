# Use the official Go image as the base image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o app main.go apiServer.go

# Set a default port number (can be overridden by user)
ENV PORT 8080

# Expose the port that the application will listen on
EXPOSE $PORT

# Define the command to run when the container starts
CMD ["./app"]

