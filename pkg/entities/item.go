package entities

import (
	gen "github.com/bersen66/wb_task0/pkg/entities/generators"
	"math/rand"
)

type Item struct {
	ChrtId      int32  `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       uint32 `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int32  `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  uint32 `json:"total_price"`
	NmId        uint32 `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      uint32 `json:"status"`
}

type Items []Item

func RandomItem(trackNumber string) Item {
	var res = Item{
		ChrtId:      rand.Int31() % 100,
		TrackNumber: trackNumber,
		Price:       rand.Uint32() % 100,
		Rid:         gen.RandomString(20, gen.DIGITS+gen.ENGLET),
		Name:        gen.RandomString(10, gen.ENGLET+gen.SPACES),
		Sale:        rand.Int31() % 100,
		Size:        gen.RandomString(4, gen.ENGLET+gen.DIGITS),
		TotalPrice:  rand.Uint32() % 100,
		NmId:        rand.Uint32() % 100,
		Brand:       gen.RandomString(10, gen.ENGLET+gen.DIGITS+gen.SPACES),
		Status:      rand.Uint32() % 100,
	}
	return res
}

func RandomItems(trackNumber string, size int) Items {
	result := make(Items, 0, size)
	for i := 0; i < size; i++ {
		result = append(result, RandomItem(trackNumber))
	}
	return result
}
