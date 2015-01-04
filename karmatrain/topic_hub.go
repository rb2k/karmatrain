package karmatrain

import (
	"github.com/Shopify/sarama"
)

type TopicHub struct {
	name           string
	ws_connections map[*connection]bool
	ws_register    chan *connection
	ws_unregister  chan *connection
	kafka_consumer *KafkaTopicConsumer
}

func NewTopicHub(topic string) (*TopicHub, error) {
	hub := &TopicHub{
		name:           topic,
		ws_connections: make(map[*connection]bool),
		ws_register:    make(chan *connection),
		ws_unregister:  make(chan *connection),
	}
	kafka_consumer, err := NewKafkaConsumer(topic)
	if err == nil {
		hub.kafka_consumer = kafka_consumer
		go hub.Run()
	}
	return hub, err
}

func (h *TopicHub) Run() {
	for {
		select {
		case c := <-h.ws_register:
			h.ws_connections[c] = true
		case c := <-h.ws_unregister:
			if _, ok := h.ws_connections[c]; ok {
				delete(h.ws_connections, c)
				close(c.send)
			}
		case event := <-h.kafka_consumer.Events():
			if event.Err != nil {
				continue
			}
			for ws := range h.ws_connections {
				ws.send <- h.kafkaEventToWsMessage(event)
			}
		}
	}
}

func (h *TopicHub) kafkaEventToWsMessage(event *sarama.ConsumerEvent) []byte {
	return event.Value
}

func (h *TopicHub) Name() string {
	return h.name
}
