package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/atoyr/virtual-arena/internal/hub"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// クロスオリジン対応
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWS(h *hub.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		roomID := c.Query("room")
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("error upgrading connection: %v", err)
			return
		}
		client := hub.NewClient(conn, roomID)
		h.Register(client)
		go client.WritePump()
		go client.ReadPump(h)
	}
}
