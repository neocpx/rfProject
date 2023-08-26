#!/bin/bash

# Build the Docker image
docker build -t server .

# Run the Docker container
docker run -p 8080:8080 server

