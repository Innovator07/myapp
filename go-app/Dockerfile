# Use the official Golang image as a base image
FROM golang:1.22.4 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy all files from the current directory to the container's working directory
COPY . .

# Build the Go app without Go modules
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM scratch

# Copy the compiled executable from the builder stage
COPY --from=builder /app/main /app/main

# Expose port 5002 to the outside world
EXPOSE 5002

# Command to run the executable
CMD ["/app/main"]

