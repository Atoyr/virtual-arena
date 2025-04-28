package hub

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"

	"github.com/atoyr/virtual-arena/internal/protocol"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type ClientInterface interface {
	Send([]byte)
}

// NewClient はクライアントのインスタンスを作成する
func NewClient(conn *websocket.Conn, roomID string) *Client {
	return &Client{
		conn:   conn,
		send:   make(chan []byte, 256),
		roomID: roomID,
	}
}

type Client struct {
	conn   *websocket.Conn
	send   chan []byte
	roomID string
}

func (c *Client) Send(message []byte) {
	select {
	case c.send <- message:
	}
}

// ReadPump はクライアントからのメッセージを読み込む
func (c *Client) ReadPump(h *Hub) {
	defer func() {
		h.Unregister(c)
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		// メッセージを処理する
		var msg protocol.Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Printf("error parsing message: %v", err)
			continue
		}

		// タイムスタンプを更新
		msg.Timestamp = time.Now().UnixNano() / int64(time.Millisecond)

		// メッセージをブロードキャスト
		updatedMsg, err := json.Marshal(msg)
		if err != nil {
			log.Printf("error marshaling message: %v", err)
			continue
		}

		h.Broadcast(updatedMsg)
	}
}

// WritePump はクライアントへメッセージを送信する
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// ハブがチャネルを閉じた
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// キューに残っているメッセージも送信
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
