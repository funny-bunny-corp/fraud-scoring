# Fraud Scoring System

A real-time fraud detection system built with Go that analyzes transaction data using multiple scoring algorithms to identify potentially fraudulent activities.

## Table of Contents
- [Overview](#overview)
- [Architecture](#architecture)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Documentation](#api-documentation)
- [Testing](#testing)
- [Development](#development)
- [Contributing](#contributing)
- [License](#license)

## Overview

The Fraud Scoring System is a microservice designed to process payment transactions in real-time and generate risk scores based on various factors including:

- Transaction value analysis
- Seller reliability scoring
- Average value comparisons
- Currency stability metrics

The system uses Apache Kafka for event streaming and gRPC for service communication, providing a scalable and efficient fraud detection solution.

## Architecture

The system follows a clean architecture pattern with the following layers:

```
├── cmd/                    # Application entry point
├── api/                    # API specifications (AsyncAPI, gRPC)
├── internal/
│   ├── domain/            # Business logic and entities
│   │   ├── application/   # Application services
│   │   ├── repositories/  # Data access interfaces
│   │   └── scoring/       # Scoring algorithms
│   ├── infra/             # Infrastructure layer
│   │   ├── grpc/         # gRPC implementations
│   │   ├── kafka/        # Kafka producers/consumers
│   │   └── logger/       # Logging configuration
│   └── adapter/          # External service adapters
│       ├── grpc/         # gRPC client adapters
│       └── kafka/        # Kafka adapters
```

### Key Components

- **Domain Layer**: Contains business logic, entities, and scoring algorithms
- **Infrastructure Layer**: Handles external communication (Kafka, gRPC)
- **Adapter Layer**: Provides interfaces to external services
- **Application Layer**: Orchestrates business operations

## Features

- **Real-time Processing**: Processes transactions as they occur via Kafka streams
- **Multi-factor Scoring**: Evaluates transactions using multiple scoring algorithms:
  - Value-based scoring
  - Seller reliability assessment
  - Historical average comparisons
  - Currency stability analysis
- **gRPC Services**: Provides high-performance RPC services for transaction data
- **Event-driven Architecture**: Uses CloudEvents for standardized event messaging
- **Scalable Design**: Built for horizontal scaling with Kafka partitioning

## Prerequisites

- Go 1.21 or later
- Apache Kafka 2.8+
- Protocol Buffers compiler (protoc)
- Docker (optional, for containerized deployment)

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd fraud-scoring
```

2. Install dependencies:
```bash
go mod download
```

3. Generate protocol buffer files:
```bash
protoc --go_out=. --go-grpc_out=. api/payment-processing.proto
```

4. Build the application:
```bash
go build -o bin/fraud-scoring cmd/main.go
```

## Configuration

The application uses environment variables for configuration:

### Basic Configuration

| Variable                         | Description                           | Default |
|----------------------------------|---------------------------------------|---------|
| `KAFKA_HOST`                     | Kafka broker host                     | localhost:9092 |
| `KAFKA_PAYMENT_PROCESSING_TOPIC` | Payment processing topic              | payment-processing |
| `KAFKA_FRAUD_DETECTION_TOPIC`    | Fraud detection topic                 | fraud-detection |
| `KAFKA_GROUP_ID`                 | Kafka consumer group ID               | fraud-scoring-group |
| `USER_TRANSACTIONS_HOST`         | User transactions service host        | localhost:8080 |

### Advanced Configuration

| Variable                | Description                    | Default |
|-------------------------|--------------------------------|---------|
| `GRPC_PORT`            | gRPC server port               | 50051   |
| `LOG_LEVEL`            | Logging level (debug/info/warn/error) | info    |
| `WORKER_COUNT`         | Number of concurrent workers   | 10      |

## Usage

### Running the Application

1. Set required environment variables:
```bash
export KAFKA_HOST=localhost:9092
export USER_TRANSACTIONS_HOST=localhost:8080
```

2. Start the application:
```bash
./bin/fraud-scoring
```

### Docker Deployment

```bash
# Build Docker image
docker build -t fraud-scoring .

# Run container
docker run -d \
  -e KAFKA_HOST=kafka:9092 \
  -e USER_TRANSACTIONS_HOST=user-service:8080 \
  fraud-scoring
```

## API Documentation

### AsyncAPI

The system implements AsyncAPI 2.6.0 specification for event-driven communication. See `api/fraud-scoring.yaml` for detailed specifications.

Key events:
- **Transaction ScoreCard Events**: Real-time scoring results
- **Payment Processing Events**: Transaction processing notifications

### gRPC Services

The system provides gRPC services for synchronous operations:

#### UserTransactionsService

- `GetUserMonthAverage`: Retrieves user's monthly transaction average
- `GetLastUserTransaction`: Gets the most recent transaction for a user

See `api/payment-processing.proto` for detailed service definitions.

## Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

### Test Structure

Tests are organized by layer:
- `internal/domain/*_test.go`: Unit tests for business logic
- `internal/infra/*_test.go`: Infrastructure layer tests
- `internal/adapter/*_test.go`: Adapter layer tests

## Development

### Project Structure

The project follows Go best practices and clean architecture principles:

- **Dependency Injection**: Uses Google Wire for dependency injection
- **Error Handling**: Structured error handling with custom error types
- **Logging**: Structured logging with Uber Zap
- **Testing**: Comprehensive unit tests with mocking

### Code Generation

The project uses code generation for:
- Protocol Buffers (gRPC services)
- Dependency injection (Wire)

Regenerate code:
```bash
make generate
```

### Adding New Features

1. Add business logic to the domain layer
2. Implement infrastructure interfaces
3. Add corresponding tests
4. Update API specifications
5. Update documentation

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.


