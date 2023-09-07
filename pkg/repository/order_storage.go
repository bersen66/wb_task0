package repository

import (
	"github.com/bersen66/wb_task0/pkg/entities"
)

type OrdersStorage interface {
	InsertOrder(order *entities.Order) error
	GetOrder(uuid string) (*entities.Order, error)
}
