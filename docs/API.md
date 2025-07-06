# API Documentation

This document provides comprehensive information about the Fraud Scoring System APIs.

## Table of Contents

- [Overview](#overview)
- [AsyncAPI (Event-Driven)](#asyncapi-event-driven)
- [OpenAPI (REST)](#openapi-rest)
- [gRPC Services](#grpc-services)
- [Authentication](#authentication)
- [Error Handling](#error-handling)
- [Rate Limiting](#rate-limiting)
- [Examples](#examples)

## Overview

The Fraud Scoring System provides three types of APIs:

1. **AsyncAPI**: Event-driven communication via Apache Kafka
2. **OpenAPI**: Synchronous REST API for HTTP clients
3. **gRPC**: High-performance RPC for service-to-service communication

## AsyncAPI (Event-Driven)

### Specification
- **File**: [`api/fraud-scoring.yaml`](../api/fraud-scoring.yaml)
- **Version**: AsyncAPI 2.6.0
- **Protocol**: Apache Kafka
- **Content Type**: `application/cloudevents+json`

### Event Channels

#### 1. Transaction Scorecard Events
```yaml
Channel: fraud-detection.transaction-scorecard
Operation: subscribe
Purpose: Receive real-time fraud scoring results
```

**Event Example:**
```json
{
  "specversion": "1.0",
  "type": "com.company.fraud-detection.transaction-scorecard.created",
  "source": "fraud-scoring-service",
  "id": "b1e0a3d7-8b2c-4f3e-9d5a-6c7b8a9d0e1f",
  "time": "2024-01-15T10:30:00Z",
  "datacontenttype": "application/json",
  "data": {
    "score": {
      "valueScore": { "score": 85, "factors": ["amount_within_normal_range"] },
      "sellerScore": { "score": 90, "factors": ["high_reputation", "verified_seller"] },
      "averageValueScore": { "score": 75, "factors": ["within_historical_range"] },
      "currencyScore": { "score": 95, "factors": ["stable_currency"] },
      "overallScore": 86,
      "riskLevel": "low"
    },
    "transaction": {
      "participants": {
        "buyer": { "document": "12345678901", "name": "John Doe" },
        "seller": { "sellerId": "seller-123" }
      },
      "order": {
        "id": "order-456",
        "at": "2024-01-15T10:25:00Z"
      },
      "payment": {
        "id": "payment-789",
        "amount": "100.00",
        "currency": "USD",
        "status": "completed"
      }
    },
    "timestamp": "2024-01-15T10:30:00Z"
  }
}
```

#### 2. Transaction Processing Events
```yaml
Channel: payment-processing.transaction-events
Operation: publish
Purpose: Send transaction data for fraud analysis
```

#### 3. Fraud Alert Events
```yaml
Channel: fraud-detection.alerts
Operation: publish
Purpose: High-priority fraud alerts
```

#### 4. User Behavior Events
```yaml
Channel: user-analytics.behavior
Operation: subscribe
Purpose: User behavior analytics for improved detection
```

### Kafka Configuration

**Producers:**
```yaml
bootstrap.servers: kafka.company.org:9092
security.protocol: SASL_SSL
sasl.mechanism: SCRAM-SHA-256
```

**Consumers:**
```yaml
group.id: fraud-detection-processors
auto.offset.reset: earliest
enable.auto.commit: false
```

## OpenAPI (REST)

### Specification
- **File**: [`api/openapi.yaml`](../api/openapi.yaml)
- **Version**: OpenAPI 3.0.3
- **Base URL**: `https://api.fraud-scoring.company.com/v1`

### Endpoints

#### Health Check
```http
GET /health
```

**Response:**
```json
{
  "status": "healthy",
  "timestamp": "2024-01-15T10:30:00Z",
  "version": "1.0.0",
  "dependencies": {
    "kafka": "healthy",
    "grpc": "healthy"
  }
}
```

#### Transaction Scoring
```http
POST /score/transaction
Content-Type: application/json
```

**Request Body:**
```json
{
  "transaction": {
    "participants": {
      "buyer": {
        "document": "12345678901",
        "name": "John Doe"
      },
      "seller": {
        "sellerId": "seller-123"
      }
    },
    "order": {
      "id": "order-456",
      "paymentType": {
        "cardInfo": "****1234",
        "token": "tok_123"
      },
      "at": "2024-01-15T10:25:00Z"
    },
    "payment": {
      "id": "payment-789",
      "amount": "100.00",
      "currency": "USD",
      "status": "completed"
    }
  }
}
```

**Response:**
```json
{
  "score": {
    "valueScore": { "score": 85, "factors": ["amount_within_normal_range"] },
    "sellerScore": { "score": 90, "factors": ["high_reputation"] },
    "averageValueScore": { "score": 75, "factors": ["within_historical_range"] },
    "currencyScore": { "score": 95, "factors": ["stable_currency"] },
    "overallScore": 86
  },
  "transaction": { /* ... */ },
  "riskLevel": "low",
  "timestamp": "2024-01-15T10:30:00Z"
}
```

#### User Analytics
```http
GET /users/{document}/transactions/average?month=2024-01
```

**Response:**
```json
{
  "document": "12345678901",
  "month": "2024-01",
  "total": "1250.75",
  "transactionCount": 15,
  "average": "83.38"
}
```

```http
GET /users/{document}/transactions/last
```

**Response:**
```json
{
  "document": "12345678901",
  "sellerId": "seller-123",
  "currency": "USD",
  "value": "100.50",
  "timestamp": "2024-01-15T10:30:00Z"
}
```

#### Metrics
```http
GET /metrics
```

**Response:**
```json
{
  "transactions": {
    "total": 10000,
    "scored": 9950,
    "failed": 50
  },
  "scoring": {
    "averageProcessingTime": 45.5,
    "riskDistribution": {
      "low": 7500,
      "medium": 2000,
      "high": 400,
      "critical": 50
    }
  },
  "system": {
    "uptime": 99.9,
    "lastRestart": "2024-01-14T08:00:00Z"
  }
}
```

## gRPC Services

### Specification
- **File**: [`api/payment-processing.proto`](../api/payment-processing.proto)
- **Package**: `user`
- **Port**: `50051`

### UserTransactionsService

#### GetUserMonthAverage
```protobuf
rpc GetUserMonthAverage(UserMonthAverageRequest) returns (UserMonthAverageResponse);
```

**Request:**
```protobuf
message UserMonthAverageRequest {
  string document = 1;
  string month = 2;
}
```

**Response:**
```protobuf
message UserMonthAverageResponse {
  string month = 1;
  string document = 2;
  string total = 3;
}
```

**Example Usage (Go):**
```go
client := user.NewUserTransactionsServiceClient(conn)
req := &user.UserMonthAverageRequest{
    Document: "12345678901",
    Month:    "2024-01",
}
resp, err := client.GetUserMonthAverage(ctx, req)
```

#### GetLastUserTransaction
```protobuf
rpc GetLastUserTransaction(LastUserTransactionRequest) returns (LastUserTransactionResponse);
```

**Request:**
```protobuf
message LastUserTransactionRequest {
  string document = 1;
}
```

**Response:**
```protobuf
message LastUserTransactionResponse {
  string document = 1;
  string sellerId = 2;
  string currency = 3;
  string value = 4;
}
```

**Example Usage (grpcurl):**
```bash
grpcurl -plaintext \
  -d '{"document": "12345678901"}' \
  localhost:50051 \
  user.UserTransactionsService/GetLastUserTransaction
```

## Authentication

### API Key Authentication
```http
X-API-Key: your-api-key-here
```

### JWT Bearer Token
```http
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### gRPC Authentication
```go
// Using interceptors
conn, err := grpc.Dial(
    "localhost:50051",
    grpc.WithTransportCredentials(insecure.NewCredentials()),
    grpc.WithPerRPCCredentials(tokenAuth{token: "your-token"}),
)
```

## Error Handling

### HTTP Error Responses
```json
{
  "error": "Invalid request parameters",
  "code": "INVALID_REQUEST",
  "timestamp": "2024-01-15T10:30:00Z",
  "details": {
    "field": "document",
    "reason": "Document number is required"
  }
}
```

### Common HTTP Status Codes
- `200 OK`: Successful request
- `400 Bad Request`: Invalid request parameters
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Resource not found
- `429 Too Many Requests`: Rate limit exceeded
- `500 Internal Server Error`: Server error
- `503 Service Unavailable`: Service temporarily unavailable

### gRPC Error Codes
- `OK`: Success
- `INVALID_ARGUMENT`: Invalid request parameters
- `UNAUTHENTICATED`: Authentication required
- `PERMISSION_DENIED`: Insufficient permissions
- `NOT_FOUND`: Resource not found
- `RESOURCE_EXHAUSTED`: Rate limit exceeded
- `INTERNAL`: Internal server error
- `UNAVAILABLE`: Service unavailable

### Kafka Error Handling
- **Producer Errors**: Delivery failures, serialization errors
- **Consumer Errors**: Deserialization errors, processing failures
- **Connection Errors**: Broker unavailable, authentication failures

## Rate Limiting

### REST API Limits
- **Per API Key**: 1000 requests/minute
- **Per IP**: 100 requests/minute
- **Burst**: Up to 10 requests/second

### gRPC Limits
- **Per Connection**: 100 RPC/second
- **Concurrent Streams**: 100 per connection

### Kafka Limits
- **Producer**: 1MB/second per partition
- **Consumer**: No explicit limits (depends on processing capacity)

## Examples

### Complete Workflow Example

#### 1. Publish Transaction Event (Kafka)
```json
{
  "specversion": "1.0",
  "type": "com.company.payment-processing.transaction.created",
  "source": "payment-gateway",
  "id": "txn-001",
  "time": "2024-01-15T10:25:00Z",
  "data": {
    "participants": {
      "buyer": { "document": "12345678901", "name": "John Doe" },
      "seller": { "sellerId": "seller-123" }
    },
    "order": { "id": "order-456", "at": "2024-01-15T10:25:00Z" },
    "payment": {
      "id": "payment-789",
      "amount": "100.00",
      "currency": "USD",
      "status": "completed"
    }
  }
}
```

#### 2. Get User History (gRPC)
```bash
grpcurl -plaintext \
  -d '{"document": "12345678901", "month": "2024-01"}' \
  localhost:50051 \
  user.UserTransactionsService/GetUserMonthAverage
```

#### 3. Receive Scorecard (Kafka)
```json
{
  "specversion": "1.0",
  "type": "com.company.fraud-detection.transaction-scorecard.created",
  "source": "fraud-scoring-service",
  "id": "score-001",
  "time": "2024-01-15T10:30:00Z",
  "data": {
    "score": { "overallScore": 86, "riskLevel": "low" },
    "transaction": { /* original transaction data */ }
  }
}
```

#### 4. Query Metrics (REST)
```bash
curl -H "X-API-Key: your-api-key" \
  https://api.fraud-scoring.company.com/v1/metrics
```

### Client Libraries

#### Go Client Example
```go
package main

import (
    "context"
    "log"
    
    "google.golang.org/grpc"
    "fraud-scoring/api/user"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", 
        grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    
    client := user.NewUserTransactionsServiceClient(conn)
    
    resp, err := client.GetUserMonthAverage(context.Background(), 
        &user.UserMonthAverageRequest{
            Document: "12345678901",
            Month:    "2024-01",
        })
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Average: %s", resp.Total)
}
```

#### JavaScript/Node.js Example
```javascript
const axios = require('axios');

async function scoreTransaction(transactionData) {
    try {
        const response = await axios.post(
            'https://api.fraud-scoring.company.com/v1/score/transaction',
            { transaction: transactionData },
            {
                headers: {
                    'X-API-Key': 'your-api-key',
                    'Content-Type': 'application/json'
                }
            }
        );
        
        return response.data;
    } catch (error) {
        console.error('Error scoring transaction:', error.response.data);
        throw error;
    }
}
```

#### Python Example
```python
import grpc
from api import user_pb2, user_pb2_grpc

def get_user_average(document, month):
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = user_pb2_grpc.UserTransactionsServiceStub(channel)
        request = user_pb2.UserMonthAverageRequest(
            document=document,
            month=month
        )
        response = stub.GetUserMonthAverage(request)
        return response
```

## Testing

### API Testing Tools
- **REST**: Postman, curl, HTTPie
- **gRPC**: grpcurl, BloomRPC, Postman
- **Kafka**: kafka-console-producer, kafka-console-consumer

### Test Environment
- **Base URL**: `https://staging-api.fraud-scoring.company.com/v1`
- **gRPC**: `staging-grpc.fraud-scoring.company.com:50051`
- **Kafka**: `staging-kafka.fraud-scoring.company.com:9092`

For more information, see the [Development Guide](DEVELOPMENT.md).