package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/atoyr/virtual-arena/internal/handlers"
	"github.com/atoyr/virtual-arena/internal/hub"
)

func main() {
	// ハブの作成と実行
	h := hub.NewHub()
	go h.Run()

	// Ginルーターの設定
	r := gin.Default()

	// 静的ファイル配信
	// FIXME: 静的ファイルのパスを修正
	r.Static("/", "../frontend/public")

	// WebSocketエンドポイント
	r.GET("/ws", handlers.ServeWS(h))

	// サーバー起動
	log.Println("backend listening on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("server error:", err)
	}
}
