package server

// websocket connection handling of server

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader websocket.Upgrader
var messageReceiver *webSocketService

func InitializeWebSocket() {
	upgrader = websocket.Upgrader{CheckOrigin: checkOriginHost}
	messageReceiver = &webSocketService{}
	messageReceiver.New()
}

// Entry Point of WebSocket Request
func Ws(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) //upgrade from http to ws
	if err != nil {
		Logger.Errorln("Upgrade to websocket failed:: " + err.Error())
		return
	}
	defer conn.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)
	// check websocket connection is kept connected.
	observeChan := connectionObserver(conn)
	// wait for client message to arrive.
	msgChan := messageListener(conn)

	// if observer or messageLister Ends, Finish WebSocket Connection.
	for {
		select {
		case <-observeChan:
			wg.Done()
		case <-msgChan:
			wg.Done()
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
	wg.Wait()
}

// check websocket is alive in interval
func connectionObserver(conn *websocket.Conn) <-chan int {
	obs := make(chan int)

	go func() {
		defer close(obs)
	observeLoop:
		for {
			var pingMessage = `server ping`
			// Do ping pong each 5 seconds to check socket is connected.
			err := conn.WriteControl(websocket.PingMessage, []byte(pingMessage), time.Now().Add(5*time.Second))
			if err != nil {
				break observeLoop
			}
		}
		obs <- 0
	}()

	return obs
}

func messageListener(conn *websocket.Conn) <-chan int {
	msgChan := make(chan int)
	//handle every client message here.
	go func() {
	messageLoop:
		for {
			mt, msg, err := conn.ReadMessage()
			if err != nil {
				Logger.Errorln(err)
				break messageLoop
			}
			switch mt {
			case websocket.BinaryMessage:
				messageReceiver.BinaryMessageReceiver(conn, msg)
			case websocket.TextMessage:
				messageReceiver.TextMessageReceiver(conn, msg)
			case websocket.PingMessage:
				messageReceiver.PingMessageReceiver(conn, msg)
			case websocket.PongMessage:
				messageReceiver.PongMessageReceiver(conn, msg)
			default:
				Logger.Errorln(`Unknown Message Type detected.`)
				break messageLoop
			}
		}
		msgChan <- 0
	}()
	return msgChan
}

// check whether websocket is connected from right host
func checkOriginHost(r *http.Request) bool {
	// for example, you can't connect this websocket server from google.
	if strings.Contains(r.Host, `google.com`) {
		return false
	}
	return true
}
