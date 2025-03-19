package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// Abilita la connessione da tutte le origini
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[string]*websocket.Conn)

func HandleConnection(writer http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println("Impossibile aggiornare la connessione: ", err)
		return
	}

	log.Println("Connessione stabilita")

	clientID := request.URL.Query().Get("id")

	if clientID == "" {
		clientID = "unknown-client"
	}

	clients[clientID] = conn

	// Stampa i client connessi
	for id := range clients {
		log.Println("Client connesso: ", id)
	}

	// Gestisce la connessione in una goroutine separata per ogni client
	go HandleClientMessages(clientID, conn) // Funzione presente in /internal/server/handlers (stessa cartella di questo file)
}
