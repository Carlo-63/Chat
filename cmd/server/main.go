package main

import (
	"chat/internal/server/handlers"
	"log"
	"net/http"
)

func main() {
	// Gestisci le richieste WebSocket
	http.HandleFunc("/ws", handlers.HandleConnection)

	// Canale per gestire il completamento dell'avvio (per non far chiudere subito il server)
	done := make(chan bool)

	// Avvia il server su porta 8080
	go func() {
		log.Println("Server WebSocket in ascolto su :8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// Attende finchè il server non si chiude
	<-done
}
