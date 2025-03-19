package utils

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

func ConnectToServer(url string) (*websocket.Conn, error) {
	// Timeout
	dialer := websocket.DefaultDialer
	dialer.HandshakeTimeout = 5 * time.Second

	// Prova a connettersi
	conn, _, err := dialer.Dial(url, nil)
	if err != nil {
		fmt.Println("Errore nella connessione al WebSocket: ", err)
		return conn, err
	}

	fmt.Println("Connessione al Websocket completata")
	return conn, nil
}
