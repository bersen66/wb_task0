package natss

import (
	"github.com/nats-io/stan.go"
	"log"
)

type SubscriberConfig struct {
	ClusterID string
	ClientID  string
}

type Subscriber struct {
	config        SubscriberConfig
	connection    stan.Conn
	subscriptions map[string]stan.Subscription
}

func mustConnect(clusterID, clientID string) stan.Conn {
	conn, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatalln(err)
	}
	return conn
}

func NewSubscriber(config SubscriberConfig) (*Subscriber, error) {
	result := new(Subscriber)

	result.config = config
	result.subscriptions = make(map[string]stan.Subscription)
	result.connection = mustConnect(config.ClusterID, config.ClientID)

	return result, nil
}

func (s *Subscriber) MustSubscribe(topic string, callback func(m *stan.Msg)) {
	subs, err := s.connection.Subscribe(topic, callback)
	if err != nil {
		log.Fatalln(err)
	}
	s.subscriptions[topic] = subs
}

func (s *Subscriber) Unsubscribe(topic string) {
	s.subscriptions[topic].Unsubscribe()
	delete(s.subscriptions, topic)
}

func (s *Subscriber) Shutdown() {
	for topic, _ := range s.subscriptions {
		s.Unsubscribe(topic)
	}
	s.connection.Close()
}
