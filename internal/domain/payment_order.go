package domain

type PaymentOrder struct {
	Id            string
	Amount        string
	Currency      string
	SellerId      string
	BuyerDocument string
}
