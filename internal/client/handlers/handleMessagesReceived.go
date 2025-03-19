package handlers

import (
	"chat/internal/types"
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

func HandleMessagesReceived(conn *websocket.Conn) {
	for {
		_, msgBytes, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("\n\nErrore nella lettura della risposta dal server")
			return
		}

		// Decodifica del messaggio
		var message types.Message

		err = json.Unmarshal(msgBytes, &message)
		if err != nil {
			fmt.Println("\n\nErrore nella conversione del messaggio: ", err)
			return
		}

		fmt.Printf("\n\n[MESSAGGIO DAL SERVER]\nContenuto: %s\nMittente: %s\n\n", message.Message, message.Sender)

		defer conn.Close()
	}
}
