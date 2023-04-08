package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sho621bake/real-time-chat-app/server/src/domain"
	"github.com/sho621bake/real-time-chat-app/server/src/handler"
	"github.com/sho621bake/real-time-chat-app/server/src/services"
)

func main() {
	pubsub := services.NewPubSubService()
	hub := domain.NewHub(pubsub)
	go hub.SubscribeMessages(pubsub)
	go hub.RunLoop()
	http.HandleFunc("/ws", handlers.NewWebsocketHandler(hub).Handle)

	port := "80"
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Panicln("Serve Error:", err)
	}
}
