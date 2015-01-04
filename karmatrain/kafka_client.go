package karmatrain

import (
	"github.com/Shopify/sarama"
)

func NewKafkaClient(clientConfig *sarama.ClientConfig) (*sarama.Client, error) {
	brokers := []string{"karmatrain.mseeger.ahserversdev.com:12333"}
	return sarama.NewClient("karamtrain", brokers, clientConfig)
}
