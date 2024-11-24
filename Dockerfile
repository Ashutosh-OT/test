# Use a minimal runtime image
FROM alpine:latest
MAINTAINER Opstree Solutions

# Set the working directory inside the container
WORKDIR /app

# Install required runtime dependencies
RUN apk add --no-cache libc6-compat bash

# Copy the prebuilt binary into the container
COPY app_binary /app/

# Ensure the binary has execute permissions
RUN chmod +x /app/app_binary

# Define the default command to run your application
ENTRYPOINT ["./app_binary"]
