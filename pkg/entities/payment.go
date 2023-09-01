package entities

import (
	gen "github.com/bersen66/wb_task0/pkg/entities/generators"
	"math/rand"
)

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestId    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       uint32 `json:"amount"`
	Payment      uint64 `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost uint32 `json:"delivery_cost"`
	GoodsTotal   uint32 `json:"goods_total"`
	CustomFee    uint32 `json:"custom_fee"`
}

func RandomPayment() Payment {
	result := Payment{
		Transaction:  gen.RandomString(10, gen.ENGLET+gen.DIGITS),
		RequestId:    gen.RandomString(rand.Int()%7, gen.ENGLET+gen.DIGITS),
		Currency:     gen.RandomCurrency(),
		Provider:     gen.RandomString(5, gen.ENGLET),
		Amount:       rand.Uint32(),
		Payment:      rand.Uint64(),
		Bank:         gen.RandomString(5, gen.ENGLET),
		DeliveryCost: rand.Uint32(),
		GoodsTotal:   rand.Uint32(),
		CustomFee:    rand.Uint32(),
	}
	return result
}
