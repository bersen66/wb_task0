package entities

type Order struct {
	Uid             string   `json:"order_uid"`
	TrackNumber     string   `json:"track_number"`
	Entry           string   `json:"entry"`
	Delivery_       Delivery `json:"delivery"`
	Payment_        Payment  `json:"payment"`
	Items           []Item   `json:"items"`
	Locale          string   `json:"locale"`
	InternalSign    string   `json:"internal_sign"`
	CustomerId      string   `json:"customer_id"`
	DeliveryService string   `json:"delivery_service"`
	ShardKey        string   `json:"shard_key"`
	SmId            uint     `json:"sm_id"`
	DateCreated     string   `json:"date_created"`
	OofShard        string   `json:"oof_shard"`
}
