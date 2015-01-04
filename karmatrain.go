package main

import (
	"fmt"
	"github.com/beejeebus/karmatrain/karmatrain"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/ws/topic/consume/{topic:[a-zA-Z0-9_-]+}", karmatrain.WebsocketApiTopicConsumer).Methods("GET")
	http.Handle("/", r)
	fmt.Println("Okay, here we go, calling http.ListenAndServe()...")
	http.ListenAndServe(":8080", nil)

	//consumer, err := karmatrain.NewKafkaConsumer("beejeebus_testing")
	//if err != nil {
	//	fmt.Println("error getting consumer")
	//	panic(err)
	//}
	//for {
	//	select {
	//	case event := <-consumer.Events():
	//		if event.Err != nil {
	//			panic(event.Err)
	//		}
	//		fmt.Printf("%s => %s\n", event.Key, event.Value)
	//	case <-time.After(60 * time.Second):
	//		fmt.Println("beeeep...")
	//	}
	//}
}