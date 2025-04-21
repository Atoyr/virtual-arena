package main

import (
    "log"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{ CheckOrigin: func(r *http.Request) bool { return true } }

func wsHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("upgrade:", err)
        return
    }
    defer conn.Close()
    // まずは PING‐PONG だけ試す
    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            log.Println("read:", err)
            break
        }
        log.Printf("recv: %s", msg)
        conn.WriteMessage(websocket.TextMessage, []byte("pong"))
    }
}

func main() {
    // 静的ファイル配信
    fs := http.FileServer(http.Dir("../frontend/public"))
    http.Handle("/", fs)
    // WebSocket エンドポイント
    http.HandleFunc("/ws", wsHandler)

    log.Println("backend listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
