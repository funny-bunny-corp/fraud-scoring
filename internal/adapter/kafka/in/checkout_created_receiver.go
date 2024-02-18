package in

import (
	"context"
	"fraud-scoring/internal/domain"
	"fraud-scoring/internal/domain/application"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"go.uber.org/zap"
)

type CheckoutEventReceiver struct {
	scr *application.PaymentRiskScoring
	log *zap.Logger
}

func (cer *CheckoutEventReceiver) Handle(ctx context.Context, event cloudevents.Event) error {
	data := &domain.CheckoutData{}
	if err := event.DataAs(data); err != nil {
		cer.log.Error("error to retrieve deserialize cloud event data", zap.String("error", err.Error()))
		return err
	}
	for _, po := range data.Payments {
		order := &domain.PaymentOrder{
			Id:            po.Id,
			Amount:        po.Amount,
			Currency:      po.Currency,
			SellerId:      po.SellerInfo.SellerId,
			BuyerDocument: data.Checkout.BuyerInfo.Document,
		}
		err := cer.scr.Assessment(order)
		if err != nil {
			cer.log.Error("error to make scorecard for transaction", zap.String("id", order.Id))
			return err
		}
	}
	return nil
}

func NewCheckoutEventReceiver(scr *application.PaymentRiskScoring, log *zap.Logger) *CheckoutEventReceiver {
	return &CheckoutEventReceiver{scr: scr, log: log}
}
