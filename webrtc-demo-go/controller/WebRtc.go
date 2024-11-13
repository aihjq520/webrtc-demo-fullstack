package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var rtcUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var connetMap = make(map[string]*websocket.Conn)
var roomMap = make(map[string]*map[string]*websocket.Conn)

func WebRtcHandler(c *gin.Context) {
	roomId := c.Query("roomId")
	userId := c.Query("userId")
	ws, err := rtcUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}
	log.Println("New Client Connected!")

	connetMap[userId] = ws

	_, rootMapOK := roomMap[roomId]
	if !rootMapOK {
		roomMap[roomId] = &connetMap
	}

	for {
		_, data, err := ws.ReadMessage()
		if err != nil {
			log.Println("Error reading from the client:", err)
		}
		connect, ok := roomMap[roomId]
		if ok {
			for _, v := range *connect {
				fmt.Println("data +%v", connect)
				v.WriteMessage(websocket.TextMessage, data)
			}
		}
		if err != nil {
			log.Println("Error reading from the client:", err)
			break
		}
	}
}
