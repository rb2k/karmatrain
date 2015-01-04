package karmatrain

import (
	"github.com/Shopify/sarama"
)

type KafkaTopicConsumer struct {
	events chan *sarama.ConsumerEvent
}

func NewKafkaConsumer(topic string) (*KafkaTopicConsumer, error) {
	client, err := NewKafkaClient(sarama.NewClientConfig())
	if err != nil {
		panic(err)
	}
	partitions, err := client.Partitions(topic)
	if err != nil {
		panic(err)
	}
	consumer := &KafkaTopicConsumer{
		events: make(chan *sarama.ConsumerEvent),
	}
	for _, partition := range partitions {
		config := sarama.NewConsumerConfig()
		config.OffsetMethod = sarama.OffsetMethodNewest
		sarama_consumer, err := sarama.NewConsumer(client, topic, partition, "karmatrain", config)
		if err != nil {
			panic(err)
		}
		go func(consumer *KafkaTopicConsumer, sarama_consumer *sarama.Consumer) {
			for {
				select {
				case event := <-sarama_consumer.Events():
					consumer.events <- event
				}
			}
		}(consumer, sarama_consumer)
	}
	return consumer, err
}

func (c *KafkaTopicConsumer) Events() chan *sarama.ConsumerEvent {
	return c.events
}
