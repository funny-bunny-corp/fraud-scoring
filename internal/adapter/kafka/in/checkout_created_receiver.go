package in

import (
	"context"
	"fraud-scoring/internal/domain"
	"fraud-scoring/internal/domain/application"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"go.uber.org/zap"
	"time"
)

const customDateFormat = "2006-01-02T15:04:05.000000"
const eventType = "paymentic.io.payment-processing.v1.payment.created"

type CheckoutEventReceiver struct {
	scr *application.PaymentRiskScoring
	log *zap.Logger
}

func (cer *CheckoutEventReceiver) Handle(ctx context.Context, event cloudevents.Event) error {
	if eventType == event.Type() {
		data := &domain.CheckoutData{}
		if err := event.DataAs(data); err != nil {
			cer.log.Error("error to retrieve deserialize cloud event data", zap.String("error", err.Error()))
			return err
		}
		t, err := time.Parse(customDateFormat, data.Checkout.At)
		if err != nil {
			cer.log.Error("error to parse date for transaction", zap.String("id", data.Payment.Id))
			return err
		}
		analysis := &domain.TransactionAnalysis{
			Participants: domain.Participants{
				Buyer: domain.BuyerInfo{
					Document: data.Checkout.BuyerInfo.Document,
					Name:     data.Checkout.BuyerInfo.Name,
				},
				Seller: domain.SellerInfo{SellerId: data.Payment.SellerInfo.SellerId},
			},
			Order: domain.Checkout{
				Id: data.Checkout.Id,
				PaymentType: domain.CardInfo{
					CardInfo: data.Checkout.CardInfo.CardInfo,
					Token:    data.Checkout.CardInfo.Token,
				},
				At: t,
			},
			Payment: domain.Payment{
				Amount:   data.Payment.Amount,
				Currency: data.Payment.Currency,
				Status:   data.Payment.Status,
				Id:       data.Payment.Id,
			},
		}
		err = cer.scr.Assessment(analysis)
		if err != nil {
			cer.log.Error("error to make scorecard for transaction", zap.String("id", analysis.Payment.Id))
			return err
		}
	}
	return nil
}

func NewCheckoutEventReceiver(scr *application.PaymentRiskScoring, log *zap.Logger) *CheckoutEventReceiver {
	return &CheckoutEventReceiver{scr: scr, log: log}
}
