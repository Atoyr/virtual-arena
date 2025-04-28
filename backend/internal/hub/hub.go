package hub

// NewHub はハブのインスタンスを作成する
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[ClientInterface]bool),
		broadcast:  make(chan []byte),
		register:   make(chan ClientInterface),
		unregister: make(chan ClientInterface),
	}
}

type Hub struct {
	clients    map[ClientInterface]bool
	broadcast  chan []byte // バイナリ／JSON混在メッセージ
	register   chan ClientInterface
	unregister chan ClientInterface
}

// Run はハブのメインループを実行する
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				client.Send(message)
			}
		}
	}
}

// Register はクライアントを登録する
func (h *Hub) Register(client ClientInterface) {
	h.register <- client
}

// Unregister はクライアントの登録を解除する
func (h *Hub) Unregister(client ClientInterface) {
	h.unregister <- client
}

// Broadcast はメッセージを全クライアントに送信する
func (h *Hub) Broadcast(message []byte) {
	h.broadcast <- message
}
