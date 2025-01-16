package coin

type CoinDto struct {
	Amount       float64 `json:"amount" binding:"required"`
	FromCurrency string  `json:"from_currency" binding:"required"`
	ToCurrency   string  `json:"to_currency" binding:"required"`
}
