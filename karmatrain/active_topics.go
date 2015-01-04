package karmatrain

type ActiveTopics struct {
	register_topic chan string
	topics         map[string]*TopicHub
}

func (at *ActiveTopics) RegisterConnectionForTopic(topic string, c *connection) error {
	if _, exists := at.topics[topic]; !exists {
		if topic_hub, err := NewTopicHub(topic); err == nil {
			at.topics[topic] = topic_hub
		} else {
			return err
		}
	}
	at.topics[topic].ws_register <- c
	return nil
}
