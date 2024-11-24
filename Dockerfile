FROM golang:latest AS build-env
MAINTAINER Opstree Solutions

# Set working directory
WORKDIR /app

# Copy all files from the current directory into the container
COPY . .

# Download dependencies and build the Go application
RUN go get -v -t ./... && \
    go build -o /app/ot-go-webapp .

# Use a minimal image for runtime
FROM alpine:latest
WORKDIR /app

# Install required dependencies
RUN apk add --no-cache libc6-compat bash

# Copy the built binary from the build stage
COPY --from=build-env /app/ot-go-webapp /app/

# Define the default command to run your application
ENTRYPOINT ["./ot-go-webapp"]
