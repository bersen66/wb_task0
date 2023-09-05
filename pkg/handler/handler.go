package handler

import (
	"encoding/json"
	"github.com/bersen66/wb_task0/pkg/entities"
	"github.com/bersen66/wb_task0/pkg/repository"
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
	json.Unmarshal(m.Data, &order)

	_, err := h.storage.Insert(&order)
	if err != nil {
		log.Fatal(err)
	}
}
