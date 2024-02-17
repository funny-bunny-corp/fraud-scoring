package in

import (
	"context"
	"fraud-scoring/internal/domain/application"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type CheckoutEventReceiver struct {
	scr *application.PaymentRiskScoring
}

func (cer *CheckoutEventReceiver) Handle(ctx context.Context, event cloudevents.Event) {

}

func NewCheckoutEventReceiver(scr *application.PaymentRiskScoring) *CheckoutEventReceiver {
	return &CheckoutEventReceiver{scr: scr}
}
