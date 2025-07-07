# Fraud Detection Scorecard Data Product

## Overview

The Fraud Detection Scorecard Data Product is a real-time analytical dataset that provides comprehensive fraud risk assessment information for payment transactions. This data product is designed following Data Mesh principles, offering a discoverable, understandable, and trustworthy source of fraud detection insights for analytical workloads.

## Data Domain & Business Context

**Domain**: Financial Services - Payment Processing & Risk Management  
**Owner**: Fraud Detection Team (fraud-detection@company.com)  
**Version**: 1.0.0

### Business Purpose

This data product serves as the primary source of truth for transaction fraud risk analysis, enabling:

- **Real-time Fraud Monitoring**: Immediate detection of high-risk transactions
- **Historical Risk Analysis**: Trend analysis of fraud patterns over time
- **Business Intelligence**: Data-driven insights for fraud prevention strategies
- **Compliance & Audit**: Detailed transaction risk assessment records

### Value Proposition

- **Real-time Processing**: Sub-second latency for fraud detection scoring
- **Multi-dimensional Scoring**: Comprehensive risk assessment across 4 key dimensions
- **Event-driven Architecture**: Reliable event sourcing from payment processing systems
- **Scalable Analytics**: Optimized for high-volume analytical queries

## Data Source & Architecture

### Event Source
- **Source System**: Fraud Scoring Service
- **Event Stream**: `fraud-detection.transaction-scorecard` (Apache Kafka)
- **Message Format**: CloudEvents v1.0 with JSON payload
- **Event Type**: `com.company.fraud-detection.transaction-scorecard.created`

### Data Flow
1. Payment transactions are processed by the payment gateway
2. Fraud scoring service analyzes transactions using ML algorithms
3. Risk scorecards are generated and published to Kafka
4. Apache Pinot ingests scorecards in real-time for analytical queries

## Schema Definition

### Primary Identifiers
| Field Name | Data Type | Description | Example |
|------------|-----------|-------------|---------|
| `transaction_id` | STRING | Unique transaction identifier | "txn_abc123" |
| `payment_id` | STRING | Payment system identifier | "pay_xyz789" |
| `order_id` | STRING | Order identifier | "ord_def456" |
| `cloudevents_id` | STRING | CloudEvent unique identifier | "ce_uuid123" |

### Transaction Participants
| Field Name | Data Type | Description | Example |
|------------|-----------|-------------|---------|
| `buyer_document` | STRING | Buyer's document identifier (CPF/CNPJ) | "12345678901" |
| `buyer_name` | STRING | Buyer's full name | "João Silva" |
| `seller_id` | STRING | Seller's unique identifier | "seller_123" |

### Payment Information
| Field Name | Data Type | Description | Example |
|------------|-----------|-------------|---------|
| `payment_amount` | BIG_DECIMAL | Transaction amount | 150.75 |
| `payment_currency` | STRING | Currency code (ISO 4217) | "BRL" |
| `payment_status` | STRING | Payment status | "completed" |
| `payment_token` | STRING | Payment method token | "tok_abcd1234" |
| `card_info_masked` | STRING | Masked card information | "****1234" |

### Risk Scoring Metrics
| Field Name | Data Type | Description | Range |
|------------|-----------|-------------|-------|
| `overall_risk_score` | INT | Overall fraud risk score | 0-100 |
| `value_score` | INT | Transaction value risk score | 0-100 |
| `seller_score` | INT | Seller reliability score | 0-100 |
| `average_value_score` | INT | Historical average comparison score | 0-100 |
| `currency_score` | INT | Currency stability score | 0-100 |
| `risk_level` | STRING | Risk classification | "low", "medium", "high", "critical" |

### Temporal Fields
| Field Name | Data Type | Description | Format |
|------------|-----------|-------------|---------|
| `event_timestamp` | TIMESTAMP | CloudEvent timestamp | Epoch milliseconds |
| `scoring_timestamp` | TIMESTAMP | When scoring was performed | Epoch milliseconds |
| `transaction_timestamp` | TIMESTAMP | Original transaction time | Epoch milliseconds |

### Metadata Fields
| Field Name | Data Type | Description | Example |
|------------|-----------|-------------|---------|
| `cloudevents_source` | STRING | Event source system | "fraud-scoring-service" |
| `cloudevents_type` | STRING | CloudEvent type | "com.company.fraud-detection.transaction-scorecard.created" |
| `cloudevents_specversion` | STRING | CloudEvents specification version | "1.0" |

## Data Quality & SLA

### Data Freshness
- **Target Latency**: < 5 seconds from transaction processing to availability
- **Maximum Latency**: < 30 seconds (99.9% percentile)
- **Monitoring**: Real-time lag monitoring via Kafka consumer lag metrics

### Data Completeness
- **Required Fields**: All schema fields marked as `notNull: true` are guaranteed
- **Missing Data**: < 0.1% of events may have optional fields missing
- **Validation**: Pre-ingestion validation ensures data integrity

### Data Accuracy
- **Scoring Accuracy**: Risk scores validated against historical fraud outcomes
- **Data Lineage**: Full traceability from source transaction to scorecard
- **Audit Trail**: CloudEvents metadata provides complete event provenance

### Retention Policy
- **Hot Data**: 90 days (immediate query access)
- **Cold Storage**: 7 years (compliance requirements)
- **Archival**: Automatic migration to long-term storage

## Usage Guidelines

### Intended Consumers
- **Fraud Analysts**: Real-time transaction monitoring and investigation
- **Data Scientists**: Model development and validation
- **Business Intelligence**: Fraud trend analysis and reporting
- **Compliance Teams**: Audit and regulatory reporting
- **Security Operations**: Alert triage and response

### Performance Considerations
- **Optimal Query Patterns**: Time-range queries with dimension filters
- **Indexing Strategy**: Inverted indexes on key dimensions, range indexes on metrics
- **Partitioning**: Data partitioned by `buyer_document` for efficient user-based queries

## Example Queries

### 1. Real-time High-Risk Transaction Monitoring
```sql
SELECT 
    transaction_id,
    buyer_document,
    seller_id,
    payment_amount,
    overall_risk_score,
    risk_level,
    event_timestamp
FROM fraud_detection_scorecard 
WHERE overall_risk_score >= 80 
    AND event_timestamp >= NOW() - INTERVAL '1' HOUR
ORDER BY overall_risk_score DESC
LIMIT 100;
```

### 2. Seller Risk Analysis
```sql
SELECT 
    seller_id,
    COUNT(*) as total_transactions,
    AVG(seller_score) as avg_seller_score,
    AVG(overall_risk_score) as avg_risk_score,
    SUM(CASE WHEN risk_level IN ('high', 'critical') THEN 1 ELSE 0 END) as high_risk_count
FROM fraud_detection_scorecard 
WHERE event_timestamp >= NOW() - INTERVAL '7' DAY
GROUP BY seller_id
HAVING COUNT(*) >= 10
ORDER BY avg_risk_score DESC;
```

### 3. Currency Risk Trends
```sql
SELECT 
    payment_currency,
    DATE_TRUNC('hour', event_timestamp) as hour_window,
    COUNT(*) as transaction_count,
    AVG(currency_score) as avg_currency_score,
    AVG(overall_risk_score) as avg_risk_score
FROM fraud_detection_scorecard 
WHERE event_timestamp >= NOW() - INTERVAL '1' DAY
GROUP BY payment_currency, DATE_TRUNC('hour', event_timestamp)
ORDER BY hour_window DESC, avg_risk_score DESC;
```

### 4. User Transaction Pattern Analysis
```sql
SELECT 
    buyer_document,
    COUNT(*) as transaction_count,
    AVG(payment_amount) as avg_amount,
    STDDEV(payment_amount) as amount_variance,
    AVG(overall_risk_score) as avg_risk_score,
    MAX(overall_risk_score) as max_risk_score
FROM fraud_detection_scorecard 
WHERE event_timestamp >= NOW() - INTERVAL '30' DAY
GROUP BY buyer_document
HAVING COUNT(*) >= 5
ORDER BY avg_risk_score DESC;
```

## Access & Support

### Data Access
- **Query Interface**: Apache Pinot SQL via REST API
- **Visualization**: Connect via JDBC/REST to BI tools (Tableau, Power BI, Grafana)
- **Programmatic Access**: REST API with JSON response format

### Support Channels
- **Technical Support**: fraud-detection-team@company.com
- **Documentation**: https://company.com/docs/fraud-detection-data-product
- **Slack Channel**: #fraud-detection-data-product
- **On-call Support**: 24/7 for critical production issues

### Getting Started
1. **Access Request**: Submit request via internal data catalog
2. **Onboarding**: Attend fraud detection domain training session
3. **Query Access**: Obtain credentials for Pinot query endpoint
4. **Monitoring**: Set up alerts for data quality and freshness

## Compliance & Security

### Data Privacy
- **PII Handling**: Buyer names and documents are pseudonymized in analytics
- **Access Controls**: Role-based access control (RBAC) enforced
- **Audit Logging**: All data access logged for compliance

### Regulatory Compliance
- **LGPD**: Brazilian data protection law compliance
- **PCI DSS**: Payment card industry security standards
- **SOX**: Sarbanes-Oxley financial reporting requirements

### Security Controls
- **Encryption**: Data encrypted at rest and in transit
- **Network Security**: VPC isolation with restricted access
- **Monitoring**: Continuous security monitoring and alerting

## Change Management

### Schema Evolution
- **Backward Compatibility**: New fields added without breaking existing queries
- **Versioning**: Semantic versioning for schema changes
- **Migration**: Automated migration support for major version changes

### Communication
- **Change Notifications**: 30-day advance notice for breaking changes
- **Release Notes**: Detailed change documentation
- **Consumer Impact**: Assessment of changes on existing consumers

## Related Data Products

- **Payment Transaction Stream**: Source transaction data
- **User Behavior Analytics**: Complementary behavioral insights
- **Merchant Risk Profiles**: Seller-focused risk assessment
- **Fraud Investigation Cases**: Downstream fraud case management

---

**Last Updated**: January 2024  
**Next Review**: April 2024  
**Document Version**: 1.0.0