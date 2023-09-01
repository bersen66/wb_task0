package entities

import (
	gen "github.com/bersen66/wb_task0/pkg/entities/generators"
	"github.com/google/uuid"
	"math/rand"
)

type Order struct {
	Uid             string   `json:"order_uid"`
	TrackNumber     string   `json:"track_number"`
	Entry           string   `json:"entry"`
	Delivery_       Delivery `json:"delivery"`
	Payment_        Payment  `json:"payment"`
	Items           Items    `json:"items"`
	Locale          string   `json:"locale"`
	InternalSign    string   `json:"internal_signature"`
	CustomerId      string   `json:"customer_id"`
	DeliveryService string   `json:"delivery_service"`
	ShardKey        string   `json:"shardkey"`
	SmId            uint32   `json:"sm_id"`
	DateCreated     string   `json:"date_created"`
	OofShard        string   `json:"oof_shard"`
}

func RandomOrder() Order {
	trackNum := gen.RandomString(8, gen.ENGLET)
	result := Order{
		Uid:             uuid.New().String(),
		TrackNumber:     trackNum,
		Entry:           gen.RandomString(5, gen.ENGLET),
		Delivery_:       RandomDelivery(),
		Payment_:        RandomPayment(),
		Items:           RandomItems(trackNum, rand.Int()%7),
		Locale:          gen.RandomLocale(),
		InternalSign:    gen.RandomString(3, gen.ENGLET),
		CustomerId:      gen.RandomString(5, gen.ENGLET),
		DeliveryService: gen.RandomString(5, gen.ENGLET),
		ShardKey:        gen.RandomString(1, gen.DIGITS),
		SmId:            rand.Uint32(),
		DateCreated:     gen.RandomDate(),
		OofShard:        gen.RandomString(3, gen.DIGITS),
	}
	return result
}
