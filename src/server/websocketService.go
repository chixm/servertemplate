package server

import (
	"github.com/gorilla/websocket"
)

// handles messages from websocket client

// Default Messaging Functions.

// when text message received
func textMessageReceiver(conn *websocket.Conn, msg []byte) {

}

// when binary message received
func binaryMessageReceiver(conn *websocket.Conn, msg []byte) {

}

// when ping message received
func pingMessageReceiver(conn *websocket.Conn, msg []byte) {

}

// when pong message received
func pongMessageReceiver(conn *websocket.Conn, msg []byte) {

}

type webSocketService struct {
	TextMessageReceiver   func(conn *websocket.Conn, msg []byte)
	BinaryMessageReceiver func(conn *websocket.Conn, msg []byte)
	PingMessageReceiver   func(conn *websocket.Conn, msg []byte)
	PongMessageReceiver   func(conn *websocket.Conn, msg []byte)
}

// setting default messaging functions
func (c *webSocketService) New() *webSocketService {
	c.TextMessageReceiver = textMessageReceiver
	c.BinaryMessageReceiver = binaryMessageReceiver
	c.PingMessageReceiver = pingMessageReceiver
	c.PongMessageReceiver = pongMessageReceiver
	return c
}
