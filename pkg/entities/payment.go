package entities

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestId    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       uint   `json:"amount"`
	Payment      uint64 `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost uint   `json:"delivery_cost"`
	GoodsTotal   uint   `json:"goods_total"`
	CustomFee    uint   `json:"custom_fee"`
}
