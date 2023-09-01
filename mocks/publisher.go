package main

import (
	"encoding/json"
	"github.com/bersen66/wb_task0/pkg/entities"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

const (
	CLusterID = "test-cluster"
	ClientID  = "publisher"
	Topic     = "orders"
)

func GenerateQueryBody() (string, error) {
	var order entities.Order = entities.RandomOrder()
	result, err := json.Marshal(order)
	return string(result), err
}

func RunPublishing(conn stan.Conn) error {
	for {
		doc, err := GenerateQueryBody()
		if err != nil {
			return err
		}
		err = conn.Publish(Topic, []byte(doc))
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 10)
	}
	return nil
}

func main() {
	sc, err := stan.Connect(CLusterID, ClientID)
	if err != nil {
		log.Fatalf("Connection error: %v\n", err)
	}

	err = RunPublishing(sc)
	if err != nil {
		log.Fatalf("Unhandled error: %v\n", err)
	}
	sc.Close()
}
