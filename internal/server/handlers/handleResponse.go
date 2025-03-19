package handlers

import (
	"chat/internal/shared"
	"chat/internal/types"
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

func HandleResponse(conn *websocket.Conn, text string, clients map[string]*websocket.Conn, message types.Message) {
	confirmation := types.Message{
		Sender:    "server",
		Recipient: message.Sender,
		Message:   text,
	}

	confirmationBytes, _ := json.Marshal(confirmation)

	if senderConn, exists := clients[message.Sender]; exists {
		err := shared.SendMessage(senderConn, confirmationBytes)
		if err != nil {
			log.Println("Errore nell'invio della conferma al mittente: ", err)
			return
		}
		log.Println("Risposta inviata al mittente")
	}
}
