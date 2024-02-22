package domain

type CheckoutData struct {
	Checkout struct {
		Id        string `json:"id"`
		BuyerInfo struct {
			Document string `json:"document"`
			Name     string `json:"name"`
		} `json:"buyerInfo"`
		CardInfo struct {
			CardInfo string `json:"cardInfo"`
			Token    string `json:"token"`
		} `json:"cardInfo"`
		IdempotencyKey string `json:"idempotencyKey"`
		At             string `json:"at"`
	} `json:"checkout"`
	Payment struct {
		Id         string `json:"id"`
		Amount     string `json:"amount"`
		Currency   string `json:"currency"`
		Status     string `json:"status"`
		SellerInfo struct {
			SellerId string `json:"sellerId"`
		} `json:"sellerInfo"`
		IdempotencyKey string `json:"idempotencyKey"`
	} `json:"payment"`
}
