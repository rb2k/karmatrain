package karmatrain

import (
	"github.com/gorilla/mux"
	"net/http"
)

var activeTopics = &ActiveTopics{
	register_topic: make(chan string),
	topics:         make(map[string]*TopicHub),
}

func WebsocketApiTopicConsumer(response http.ResponseWriter, request *http.Request) {
	c, err := NewWebsocket(response, request)
	if err != nil {
		return
	}
	activeTopics.RegisterConnectionForTopic(mux.Vars(request)["topic"], c)
}
