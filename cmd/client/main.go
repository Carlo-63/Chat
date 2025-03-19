package main

import (
	"bufio"
	"chat/internal/client/handlers"
	"chat/internal/client/utils"
	"chat/internal/shared"
	"chat/internal/types"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

func readInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func sendMessage(clientID string, conn *websocket.Conn) {
	msg := readInput("\nInserire il Messaggio: ")
	recipient := readInput("Inserire il Destinatario: ")

	// Crea l'oggetto da inviare
	message := types.Message{
		Sender:    clientID,
		Recipient: recipient,
		Message:   msg,
	}

	// Converte l'oggetto in JSON
	messageBytes, _ := json.Marshal(message)

	err := shared.SendMessage(conn, messageBytes)
	if err != nil {
		fmt.Println("Errore nell'invio del messaggio: ", err)
		return
	}
	fmt.Println("Messaggio inviato con successo")
}

func main() {
	var clientID string

	fmt.Print("Inserire il Client ID: ")
	fmt.Scanln(&clientID)

	url := "ws://localhost:8080/ws?id=" + clientID

	// Connessione
	conn, err := utils.ConnectToServer(url)
	if err != nil {
		fmt.Println("Impossibile connettersi al server")
		return
	}

	var ricezione bool

	for {
		fmt.Println("\nOpzioni disponibili:")
		fmt.Println("1) Invio messaggio")
		if !ricezione {
			fmt.Println("2) Avvia ricezione messaggi")
		}
		fmt.Println("9) Uscita")

		opzione := readInput("Selezionare un'opzione: ")

		switch opzione {
		case "1":
			sendMessage(clientID, conn)
		case "2":
			go func() {
				handlers.HandleMessagesReceived(conn)
			}()
			ricezione = true
		case "9":
			return
		default:
			fmt.Println("Opzione non valida")
		}
	}
}
