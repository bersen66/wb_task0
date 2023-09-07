package main

import (
	"encoding/json"
	"fmt"
	"github.com/bersen66/wb_task0/pkg/entities"
	"github.com/bersen66/wb_task0/pkg/repository"
	"log"
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
	err := storage.InsertOrder(&order)
	if err != nil {
		t.Fatalf("order: \n %v\n, err: %v", string(data), err)
	}
}

func TestGetOrder(t *testing.T) {
	storage := repository.NewPGStorage(repository.PGStorageConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "aboba",
		DBName:   "wbs0",
		SSLMode:  "disable",
	})

	//order := entities.RandomOrder()
	//_, err := storage.InsertOrder(&order)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//time.Sleep(time.Second * 10)
	fetched, err := storage.GetOrder("bbb3bfb7-68a6-44cf-98d5-032dfa070be9")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	data, err := json.Marshal(fetched)
	fmt.Println(string(data))
}
