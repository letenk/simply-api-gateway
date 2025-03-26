# Simply API Gateway

This repository demonstrates a simple API Gateway using [Fiber](https://gofiber.io/). It serves as a learning resource for understanding API Gateway concepts, as discussed in the related blog post.

## Features
- Reverse proxy for multiple backend services
- Request rate limiting
- Logging middleware

## Prerequisites
- Go 1.18+
- Make

## Installation
```sh
git clone https://github.com/letenk/simply-api-gateway.git
cd simply-api-gateway
go mod tidy  # Install missing dependencies
```

## Running the API Gateway
Use the Makefile to start all services easily:
```sh
make run  # Starts API Gateway and backend services
```

## Stopping Services
To stop all running services:
```sh
make stop
```
