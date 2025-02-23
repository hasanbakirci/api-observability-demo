# System Performance Monitoring with Prometheus, Grafana, and Nginx  

This is an example project demonstrating how to monitor system performance using Prometheus, Grafana, and Nginx. The project includes a Go-based API, RabbitMQ for messaging, and Nginx as a reverse proxy, with Prometheus collecting metrics and Grafana visualizing them.  

For a detailed explanation, please visit: [[Here](https://hasanbakirci.medium.com/sistem-performans%C4%B1n%C4%B1-i%CC%87zleme-prometheus-grafana-ve-nginx-ile-%C3%B6rnek-proje-e1cfaa5f2c60)]

## Tech Stack

- **Language**: Go
- **Web Framework**: Echo
- **Message Queue**: RabbitMQ
- **Container**: Docker
- **Metrics**: Prometheus
- **Monitoring**: Grafana
- **Reverse Proxy**: Nginx

## Getting Started

1. **Clone the repository:**

    ```bash
    git clone https://github.com/hasanbakirci/api-observability-demo
    ```

2. **Building the Docker Image:**

    ```bash
    docker build -t api-observability-demo .
    ```
2. **Start the required services using Docker Compose:**

    ```bash
    cd ./deploy
    docker-compose up -d
    ```

3. **API**
    ```bash
    go run main.go api
    ```
4. **Consumer**
    ```bash
    go run main.go consumer
    ```

## Endpoints

- `GET  /metrics`
- `GET  /health`
- `POST /api/events`
