package kafka

import "os"

type SaramaConfig struct {
	Host                   string
	PaymentProcessingTopic string
	FraudDetectionTopic    string
	GroupId                string
}

func NewSaramaConfig() *SaramaConfig {
	return &SaramaConfig{
		Host:                   os.Getenv("KAFKA_HOST"),
		PaymentProcessingTopic: os.Getenv("KAFKA_PAYMENT_PROCESSING_TOPIC"),
		FraudDetectionTopic:    os.Getenv("KAFKA_FRAUD_DETECTION_TOPIC"),
		GroupId:                os.Getenv("KAFKA_GROUP_ID"),
	}
}
