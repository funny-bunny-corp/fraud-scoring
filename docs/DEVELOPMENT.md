# Development Guide

This document provides comprehensive information for developers working on the Fraud Scoring System.

## Table of Contents

- [Architecture Overview](#architecture-overview)
- [API Specifications](#api-specifications)
- [Testing Strategy](#testing-strategy)
- [Development Workflow](#development-workflow)
- [Code Style Guidelines](#code-style-guidelines)
- [Performance Considerations](#performance-considerations)
- [Security Best Practices](#security-best-practices)
- [Troubleshooting](#troubleshooting)

## Architecture Overview

The Fraud Scoring System follows a clean architecture pattern with clear separation of concerns:

```
fraud-scoring/
├── cmd/                    # Application entry points
│   ├── main.go            # Main application
│   ├── manager.go         # Application manager
│   ├── wire.go            # Dependency injection configuration
│   └── wire_gen.go        # Generated dependency injection code
├── internal/              # Private application code
│   ├── domain/           # Business logic layer
│   │   ├── application/  # Application services
│   │   ├── repositories/ # Data access interfaces
│   │   └── scoring/      # Scoring algorithms
│   ├── infra/           # Infrastructure layer
│   │   ├── grpc/        # gRPC implementations
│   │   ├── kafka/       # Kafka producers/consumers
│   │   └── logger/      # Logging configuration
│   └── adapter/         # External service adapters
│       ├── grpc/        # gRPC client adapters
│       └── kafka/       # Kafka adapters
├── api/                  # API specifications
│   ├── fraud-scoring.yaml # AsyncAPI specification
│   ├── openapi.yaml      # OpenAPI/REST specification
│   └── *.proto          # Protocol buffer definitions
├── tests/               # Test files
│   ├── fixtures/        # Test data fixtures
│   ├── integration/     # Integration tests
│   └── e2e/            # End-to-end tests
└── docs/               # Documentation
```

### Layer Responsibilities

#### Domain Layer
- **Entities**: Core business objects (Transaction, ScoreCard, etc.)
- **Value Objects**: Immutable objects representing business concepts
- **Domain Services**: Business logic that doesn't belong to a specific entity
- **Repositories**: Interfaces for data access
- **Application Services**: Orchestrate business operations

#### Infrastructure Layer
- **Database**: Data persistence implementations
- **Message Queue**: Kafka producers and consumers
- **gRPC**: External service communication
- **Logging**: Structured logging with Zap
- **Monitoring**: Metrics and health checks

#### Adapter Layer
- **HTTP Handlers**: REST API endpoints
- **gRPC Servers**: gRPC service implementations
- **Message Handlers**: Kafka message processing
- **External Clients**: Third-party service integrations

## API Specifications

### AsyncAPI (Event-Driven Architecture)

The system uses AsyncAPI 2.6.0 for event-driven communication:

#### Key Events

1. **Transaction Processing Events**
   - Channel: `payment-processing.transaction-events`
   - Purpose: Trigger fraud analysis
   - Schema: CloudEvent with transaction data

2. **Transaction Scorecard Events**
   - Channel: `fraud-detection.transaction-scorecard`
   - Purpose: Fraud scoring results
   - Schema: CloudEvent with scorecard data

3. **Fraud Alert Events**
   - Channel: `fraud-detection.alerts`
   - Purpose: High-risk transaction alerts
   - Schema: CloudEvent with alert data

4. **User Behavior Events**
   - Channel: `user-analytics.behavior`
   - Purpose: User behavior analytics
   - Schema: CloudEvent with behavior data

#### Event Schema Example

```json
{
  "specversion": "1.0",
  "type": "com.company.fraud-detection.transaction-scorecard.created",
  "source": "fraud-scoring-service",
  "id": "uuid",
  "time": "2024-01-15T10:30:00Z",
  "datacontenttype": "application/json",
  "data": {
    "score": {
      "valueScore": { "score": 85 },
      "sellerScore": { "score": 90 },
      "averageValueScore": { "score": 75 },
      "currencyScore": { "score": 95 },
      "overallScore": 86,
      "riskLevel": "low"
    },
    "transaction": { ... },
    "timestamp": "2024-01-15T10:30:00Z"
  }
}
```

### OpenAPI (REST API)

The system provides REST endpoints for synchronous operations:

#### Key Endpoints

1. **Health Check**
   - `GET /health`
   - Returns service health status

2. **Transaction Scoring**
   - `POST /score/transaction`
   - Analyze transaction and return fraud score

3. **User Analytics**
   - `GET /users/{document}/transactions/average`
   - `GET /users/{document}/transactions/last`

4. **Metrics**
   - `GET /metrics`
   - Service operational metrics

### gRPC Services

Protocol Buffer definitions for high-performance RPC:

#### UserTransactionsService

```protobuf
service UserTransactionsService {
  rpc GetUserMonthAverage(UserMonthAverageRequest) returns (UserMonthAverageResponse);
  rpc GetLastUserTransaction(LastUserTransactionRequest) returns (LastUserTransactionResponse);
}
```

## Testing Strategy

### Test Pyramid

The testing strategy follows the test pyramid principle:

```
    /\
   /  \     E2E Tests (Few)
  /____\    
 /      \   Integration Tests (Some)
/________\  Unit Tests (Many)
```

### Test Types

#### Unit Tests
- **Location**: `internal/*/\*_test.go`
- **Purpose**: Test individual components in isolation
- **Tools**: Go testing package, mocks
- **Coverage Target**: 85%

#### Integration Tests
- **Location**: `tests/integration/`
- **Purpose**: Test component interactions
- **Tools**: Docker Compose, testcontainers
- **Scope**: Database, Kafka, gRPC integrations

#### End-to-End Tests
- **Location**: `tests/e2e/`
- **Purpose**: Test complete workflows
- **Tools**: Docker Compose, API clients
- **Scope**: Full system behavior

#### Performance Tests
- **Location**: `tests/performance/`
- **Purpose**: Benchmark performance
- **Tools**: Go benchmarks, load testing
- **Metrics**: Throughput, latency, resource usage

### Test Configuration

Test configuration is managed through `test_config.yaml`:

```yaml
environments:
  unit:
    description: "Unit test environment with mocked dependencies"
    # ... configuration
  integration:
    description: "Integration test environment with real services"
    # ... configuration
```

### Running Tests

```bash
# Unit tests
make test

# Unit tests with coverage
make test-coverage

# Integration tests
make test-integration

# E2E tests
make test-e2e

# Performance tests
make bench

# All tests
make ci
```

## Development Workflow

### 1. Setup Development Environment

```bash
# Clone repository
git clone <repository-url>
cd fraud-scoring

# Install dependencies
make deps

# Install development tools
make install-tools

# Generate code
make generate

# Run tests
make test
```

### 2. Feature Development

1. **Create Feature Branch**
   ```bash
   git checkout -b feature/new-scoring-algorithm
   ```

2. **Implement Feature**
   - Write failing tests first (TDD)
   - Implement feature
   - Ensure tests pass

3. **Run Pre-commit Checks**
   ```bash
   make pre-commit
   ```

4. **Submit Pull Request**
   - Ensure CI passes
   - Request code review
   - Address feedback

### 3. Code Review Process

- **Automated Checks**: CI pipeline runs automatically
- **Manual Review**: Team member reviews code
- **Approval**: Two approvals required for merge
- **Merge**: Squash and merge to main branch

## Code Style Guidelines

### Go Code Style

- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use [golangci-lint](https://golangci-lint.run/) for linting
- Run `make fmt` before committing
- Use meaningful variable and function names

### Package Structure

```go
// Package declaration
package domain

// Imports grouped by: standard library, third party, internal
import (
    "context"
    "time"
    
    "github.com/google/uuid"
    "go.uber.org/zap"
    
    "fraud-scoring/internal/domain/repositories"
)

// Types, constants, variables
type ScoreCard struct {
    ValueScore int `json:"valueScore"`
    // ...
}

// Functions
func NewScoreCard() *ScoreCard {
    return &ScoreCard{}
}
```

### Error Handling

```go
// Custom error types
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error on field %s: %s", e.Field, e.Message)
}

// Error handling
result, err := someOperation()
if err != nil {
    return nil, fmt.Errorf("failed to perform operation: %w", err)
}
```

### Testing Style

```go
func TestScoringService_Assessment(t *testing.T) {
    tests := []struct {
        name     string
        input    TransactionAnalysis
        expected ScoringResult
        wantErr  bool
    }{
        {
            name: "valid transaction",
            input: TransactionAnalysis{
                // test data
            },
            expected: ScoringResult{
                // expected result
            },
            wantErr: false,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test implementation
        })
    }
}
```

## Performance Considerations

### Scoring Algorithm Performance

- **Concurrent Processing**: Use goroutines for parallel scoring
- **Caching**: Cache frequently accessed data
- **Database Optimization**: Use proper indexing and query optimization
- **Resource Limits**: Set appropriate timeouts and limits

### Kafka Performance

- **Partitioning**: Use user document for partitioning
- **Batch Processing**: Process messages in batches
- **Consumer Groups**: Use multiple consumers for scalability
- **Acknowledgment**: Use manual acknowledgment for reliability

### gRPC Performance

- **Connection Pooling**: Reuse connections
- **Streaming**: Use streaming for large datasets
- **Compression**: Enable compression for large payloads
- **Timeouts**: Set appropriate timeouts

## Security Best Practices

### Authentication and Authorization

- **API Keys**: Use API keys for service-to-service communication
- **JWT Tokens**: Use JWT for user authentication
- **Role-based Access**: Implement proper authorization
- **Rate Limiting**: Prevent abuse with rate limiting

### Data Protection

- **Encryption**: Encrypt sensitive data at rest and in transit
- **PII Handling**: Properly handle personally identifiable information
- **Data Masking**: Mask sensitive data in logs and responses
- **Audit Logging**: Log all access to sensitive data

### Infrastructure Security

- **Network Security**: Use VPCs and security groups
- **Container Security**: Scan containers for vulnerabilities
- **Secret Management**: Use proper secret management
- **Monitoring**: Monitor for security threats

## Troubleshooting

### Common Issues

#### 1. Kafka Connection Issues

```bash
# Check Kafka connectivity
kafka-topics --bootstrap-server localhost:9092 --list

# Check consumer groups
kafka-consumer-groups --bootstrap-server localhost:9092 --list
```

#### 2. gRPC Connection Issues

```bash
# Test gRPC connection
grpcurl -plaintext localhost:50051 list

# Test specific service
grpcurl -plaintext localhost:50051 user.UserTransactionsService/GetUserMonthAverage
```

#### 3. Database Connection Issues

```bash
# Check database connectivity
psql -h localhost -U user -d fraud_scoring

# Check table structure
\d+ transactions
```

### Debug Logging

Enable debug logging for troubleshooting:

```bash
export LOG_LEVEL=debug
./bin/fraud-scoring
```

### Health Checks

Check service health:

```bash
curl http://localhost:8080/health
```

### Metrics

Monitor service metrics:

```bash
curl http://localhost:8080/metrics
```

## Additional Resources

- [AsyncAPI Documentation](https://www.asyncapi.com/)
- [OpenAPI Documentation](https://swagger.io/specification/)
- [gRPC Documentation](https://grpc.io/docs/)
- [Kafka Documentation](https://kafka.apache.org/documentation/)
- [Go Testing Documentation](https://golang.org/pkg/testing/)

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

For more information, see [CONTRIBUTING.md](CONTRIBUTING.md).