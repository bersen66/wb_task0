package main

import (
	"encoding/json"
	"github.com/bersen66/wb_task0/pkg/entities"
	"github.com/bersen66/wb_task0/pkg/repository"
	"testing"
)

func TestConnection(t *testing.T) {
	storage := repository.NewPGStorage(repository.PGStorageConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "aboba",
		DBName:   "wbs0",
		SSLMode:  "disable",
	})

	order := entities.RandomOrder()
	data, _ := json.Marshal(order)
	_, err := storage.Insert(&order)
	if err != nil {
		t.Fatalf("order: \n %v\n, err: %v", string(data), err)
	}
}
