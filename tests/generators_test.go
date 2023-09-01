package main

import (
	"encoding/json"
	"fmt"
	"github.com/bersen66/wb_task0/pkg/entities"
	gen "github.com/bersen66/wb_task0/pkg/entities/generators"
	"math/rand"
	"testing"
)

// Just checking that functions works

func TestGenerateRandomString(t *testing.T) {
	fmt.Println("RANDOM STRINGS GENERATION:")
	for i := 0; i < 10; i++ {
		fmt.Println(gen.RandomString(uint(rand.Int31n(5))+1, "qazwsxedcrfvtgbyhnujmik,ol.p;"))
	}
}

func TestGeneratePhoneNumber(t *testing.T) {
	fmt.Println("RANDOM PHONE NUMBERS GENERATION:")
	for i := 0; i < 10; i++ {
		fmt.Println(gen.RandomPhoneNumber())
	}
}

func TestRandomCurrencies(t *testing.T) {
	fmt.Println("RANDOM CURRENCIES GENERATION")
	for i := 0; i < 10; i++ {
		fmt.Println(gen.RandomCurrency())
	}
}

func TestRandomDate(t *testing.T) {
	fmt.Println("RANDOM DATE GENERATION")
	for i := 0; i < 10; i++ {
		fmt.Println(gen.RandomDate())
	}
}

func TestRandomDelivery(t *testing.T) {
	fmt.Println("RANDOM DELIVERIES GENERATION")
	for i := 0; i < 10; i++ {
		dlv := entities.RandomDelivery()
		data, _ := json.Marshal(dlv)
		fmt.Println(string(data))
	}
}

func TestRandomItem(t *testing.T) {
	fmt.Println("RANDOM ITEM GENERATION")
	for i := 0; i < 10; i++ {
		item := entities.RandomItem(gen.RandomString(4, gen.ENGLET))
		data, _ := json.Marshal(item)
		fmt.Println(string(data))
	}
}

func TestRandomItems(t *testing.T) {
	fmt.Println("RANDOM ITEMS GENERATION")
	items := entities.RandomItems(gen.RandomString(3, gen.ENGLET), 10)
	data, _ := json.Marshal(items)
	fmt.Println(string(data))
}
