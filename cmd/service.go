package main

import (
	"fmt"
	"github.com/bersen66/wb_task0/pkg/handler"
	"github.com/bersen66/wb_task0/pkg/repository"
	"github.com/bersen66/wb_task0/pkg/transport/natss"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	ClusterID     = "test-cluster"
	ClientID      = "subscriber"
	CacheCapacity = 1000
	DumpFile      = "dump.txt"
)

func PollOsSignals() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	acceptedSignal := <-signalChannel
	if acceptedSignal == syscall.SIGINT {
		fmt.Println("Bye!")
	}
}

func main() {

	storage := repository.NewPGStorage(repository.PGStorageConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "aboba",
		DBName:   "wbs0",
		SSLMode:  "disable",
	})
	cache := repository.NewStorageCache(CacheCapacity, storage)
	err := cache.Load(DumpFile)
	defer cache.Dump(DumpFile)
	if err != nil {
		log.Fatal(err)
	}

	handlers := handler.NewHandler(cache)

	sb, err := natss.NewSubscriber(natss.SubscriberConfig{
		ClusterID: ClusterID,
		ClientID:  ClientID,
	})

	if err != nil {
		log.Fatal(err)
	}

	sb.MustSubscribe("orders", handlers.InsertOrder)
	defer sb.Shutdown()

	PollOsSignals()
}
