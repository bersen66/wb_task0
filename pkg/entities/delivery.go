package entities

import (
	gen "github.com/bersen66/wb_task0/pkg/entities/generators"
)

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

func RandomDelivery() Delivery {
	result := Delivery{
		Name:    gen.RandomString(10, gen.ENGLET+gen.DIGITS),
		Phone:   gen.RandomPhoneNumber(),
		Zip:     gen.RandomString(7, gen.DIGITS),
		City:    gen.RandomString(5, gen.ENGLET),
		Address: gen.RandomString(8, gen.ENGLET+gen.SPACES),
		Region:  gen.RandomString(5, gen.ENGLET+gen.SPACES),
		Email:   gen.RandomString(10, gen.ENGLET+gen.SPACES),
	}
	return result
}
