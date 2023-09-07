package handler

import (
	"encoding/json"
	"github.com/bersen66/wb_task0/pkg/entities"
	"github.com/bersen66/wb_task0/pkg/repository"
	"github.com/google/go-cmp/cmp"
	"github.com/nats-io/stan.go"
	"log"
)

type Handler struct {
	storage repository.OrdersStorage
}

func NewHandler(storage repository.OrdersStorage) *Handler {
	result := new(Handler)

	result.storage = storage

	return result
}

func (h *Handler) InsertOrder(m *stan.Msg) {
	var order = entities.Order{}

	err := json.Unmarshal(m.Data, &order)
	if err != nil {
		log.Println(err)
		return
	}
	err = h.storage.InsertOrder(&order)
	if err != nil {
		log.Fatal(err)
	}

	data, err := h.storage.GetOrder(order.Uid)
	if !cmp.Equal(order.Uid, data.Uid) {
		log.Fatal("Not eq")
	}
}
