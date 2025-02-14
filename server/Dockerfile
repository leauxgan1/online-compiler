# Use Go 1.23 bookworm as base image
FROM golang:1.23-alpine AS base

# Move to working directory /build
WORKDIR /server

# Copy the go.mod and go.sum files to the /build directory
COPY go.mod ./ 

# Install dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the application
RUN go build -o online-compiler-server .

# Document the port that may need to be published
EXPOSE 8000

# Start the application
CMD ["./online-compiler-server"]

