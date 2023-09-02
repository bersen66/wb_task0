package repository

import (
	"github.com/bersen66/wb_task0/pkg/entities"
	"github.com/google/uuid"
)

type OrdersStorage interface {
	Insert(order *entities.Order) (bool, error)
	Contains(order *entities.Order) (bool, error)
	GetOrder(uuid uuid.UUID) (entities.Order, error)
}
