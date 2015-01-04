package karmatrain

import (
	"github.com/Shopify/sarama"
)

func NewKafkaProducer() (*sarama.Producer, error) {
	client, _ := NewKafkaClient(sarama.NewClientConfig())
	return sarama.NewProducer(client, sarama.NewProducerConfig())
}
