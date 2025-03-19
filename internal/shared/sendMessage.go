package shared

import (
	"github.com/gorilla/websocket"
)

func SendMessage(conn *websocket.Conn, messageBytes []byte) error {
	err := conn.WriteMessage(websocket.TextMessage, messageBytes)
	if err != nil {
		return err
	}
	return nil
}
