
# Tracing

This repository demonstrates how to integrate Jaeger for distributed tracing in a Go-based service. It includes a basic REST API implemented in Go that sends traces to a Jaeger instance, allowing developers to visualize how requests traverse through the application.

## Prerequisites

Before you start, ensure you have the following installed:

-   Go (version 1.20 or later)
-   Docker
-   Any standard Go IDE or editor of your choice

## Quick Start

Follow these steps to get the project up and running on your local machine:

### 1. Clone the Repository

`git clone https://github.com/Salaton/tracing.git`

`cd tracing` 

### 2. Run Jaeger

Start a local Jaeger instance using Docker:


```bash 
docker run --rm --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 4317:4317 \
  -p 4318:4318 \
  -p 14250:14250 \
  -p 14268:14268 \
  -p 14269:14269 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.56
 ``` 

### 3. Build and Run the Go Application

`go build ./tracing` 

### 4. Access the Jaeger UI

Open your web browser and navigate to `http://localhost:16686` to view the Jaeger UI and the traces collected from the application.

## Features

-   Simple REST API in Go
-   Jaeger tracing integration
-   Visualization of tracing data through Jaeger UI

## Contributing

Contributions are welcome! Feel free to submit pull requests or open issues to discuss improvements or additions.
- TODO: Add GraphQL, GORM tracing