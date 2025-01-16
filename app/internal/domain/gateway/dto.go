package gateway

type CardDetails struct {
	Number string `json:"number" binding:"required"`
	Expiry string `json:"expiry" binding:"required"`
	CVV    string `json:"cvv" binding:"required"`
}

type GatewayDto struct {
	Gateway       string      `json:"gateway" binding:"required"`
	Amount        float64     `json:"amount" binding:"required"`
	Currency      string      `json:"currency" binding:"required"`
	PaymentMethod string      `json:"payment_method" binding:"required"`
	CardDetails   CardDetails `json:"card_details" binding:"required"`
}
