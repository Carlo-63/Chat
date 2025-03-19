package handlers

import (
	"chat/internal/shared"
	"chat/internal/types"
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

func HandleClientMessages(clientID string, conn *websocket.Conn) {
	defer func() {
		conn.Close()
		delete(clients, clientID)
		log.Println("Client disconnesso: ", clientID)
	}()

	// Gestione dei messaggi ricevuti
	for {
		_, msgBytes, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("Errore nella lettura del messaggio: ", err)
			}
			break
		}

		// Decodifica del messaggio
		var message types.Message

		err = json.Unmarshal(msgBytes, &message)
		if err != nil {
			log.Println("Errore nella conversione del messaggio: ", err)
			return
		}

		log.Println("Messaggio ricevuto da ", message.Sender)

		// Invia il messaggio al destinatario
		if recipientConn, exists := clients[message.Recipient]; exists {
			// Se il destinatario esiste svolge queste operazioni
			HandleResponse(conn, "Messaggio ricevuto correttamente", clients, message)

			err = shared.SendMessage(recipientConn, msgBytes)
			if err != nil {
				log.Println("Errore nell'invio della messaggio al client destinatario")
				return
			}
			log.Println("Messaggio inviato a: ", message.Recipient)
		} else {
			HandleResponse(conn, "Destinatario inesistente", clients, message)
		}

	}
}
