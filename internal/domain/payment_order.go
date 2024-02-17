package domain

import "time"

type PaymentOrder struct {
	Id            string
	Amount        string
	Currency      string
	SellerId      string
	BuyerDocument string
	At            time.Month
}
