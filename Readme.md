# UUID Generation Server

A lightweight HTTP server written in Go that generates and returns UUIDs. Designed for simplicity and reliability in distributed systems.

## Features

- Generates version 4 UUIDs (random)
- Simple REST API endpoint
- Docker container support
- CI/CD pipeline with GitHub Actions
- Nix development environment
- Linting and testing integration

## Prerequisites

- Go 1.24+
- Make
- Docker 20.10+
- Nix (optional for development environment)

## Installation

```bash
git clone https://github.com/your-org/uuid-server.git
cd uuid-server
make build
```

Binary will be created in `build/uuidserver`

## Usage

### Running the server
```bash
make run
```

### Endpoints
**Generate UUID**
```bash
curl http://localhost:8080/uuid
```
Response:
```json
{"uuid":"f47ac10b-58cc-4372-a567-0e02b2c3d479"}
```

## Docker Deployment

### Build image
```bash
docker build -t uuid-server:latest .
```

### Run container
```bash
docker run -p 8080:8080 uuid-server:latest
```

### Pull from Yandex Container Registry
```bash
docker pull cr.yandex/crpe51e4v7vs5p6ahf5m/go-uuid:latest
```

## Development

### Nix Environment
```bash
nix develop
```

### CI/CD Pipeline
Automatically builds and pushes Docker images to Yandex Container Registry on:
- New GitHub releases
- Pushes to main branch
- Tags matching semantic versioning patterns

## License
[MIT](LICENSE) - See repository for details
