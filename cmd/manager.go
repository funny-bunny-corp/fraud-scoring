package main

import (
	"context"
	"fraud-scoring/internal/adapter/kafka/in"
	"fraud-scoring/internal/infra/kafka"
)

type Manager struct {
	receiver *in.CheckoutEventReceiver
	cli      kafka.CloudEventsReceiver
}

func (m *Manager) Start() error {
	err := m.cli.StartReceiver(context.Background(), m.receiver.Handle)
	if err != nil {
		return err
	}
	return nil
}

func NewManager(receiver *in.CheckoutEventReceiver, cli kafka.CloudEventsReceiver) *Manager {
	return &Manager{
		receiver: receiver,
		cli:      cli,
	}
}
