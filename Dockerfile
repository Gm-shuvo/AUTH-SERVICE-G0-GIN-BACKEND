#### Without Hot Reloading ####

# # Start from the official Golang base image
# FROM golang:1.22.1-alpine3.18 as builder

# # Set the Current Working Directory inside the container
# WORKDIR /app

# # Copy go mod and sum files
# COPY go.mod go.sum ./

# # Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# RUN go mod download

# # Copy the source from the current directory to the Working Directory inside the container
# COPY . .

# # Build the Go app
# RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# ######## Start a new stage from scratch #######
# FROM alpine:latest  

# RUN apk --no-cache add ca-certificates

# WORKDIR /root/

# # Copy the Pre-built binary file and the .env file from the previous stage
# COPY --from=builder /app/main .
# COPY --from=builder /app/.env .

# # Command to run the executable
# CMD ["./main"]



#### With Hot Reloading ####

# Use the official Golang image to create a build artifact.
FROM golang:1.22.1-alpine3.18 as builder

# Install air for live reloading
RUN go install github.com/cosmtrek/air@latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o ./.bin/main ./cmd/main.go

# Start a new stage from scratch
FROM golang:1.22.1-alpine3.18

WORKDIR /app

COPY --from=builder /app ./
COPY --from=builder /go/bin/air /usr/local/bin/air

# Copy .air.toml configuration file for air
COPY .air.toml ./

# Command to run the application using air for hot reloading
CMD ["air", "-c", ".air.toml"]

