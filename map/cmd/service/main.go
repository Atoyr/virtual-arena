package main

import (
    "path/filepath"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/atoyr/virtual-arena/map-service/internal/handler"
)

func main() {
    r := gin.Default()

    // マップ一覧 / 編集用 API
    r.GET  ("/maps",           handler.ListMaps)
    r.POST ("/maps",           handler.CreateMap)
    r.GET  ("/maps/:id",       handler.GetMap)
    r.PUT  ("/maps/:id",       handler.UpdateMap)
    r.DELETE("/maps/:id",      handler.DeleteMap)

    // タイル静的配信（例: /maps/123/tiles/0/1/2.png）
    // tiles を動的に返すハンドラ
    r.GET("/maps/:id/tiles/*filepath", func(c *gin.Context) {
        id := c.Param("id")
        fp := c.Param("filepath")                // "/z/x/y.png" の形式
        fullPath := filepath.Join("data/maps", id, "tiles", fp)
        // 存在チェックしておくとより安全
        if _, err := filepath.Abs(fullPath); err != nil {
            c.Status(http.StatusBadRequest)
            return
        }
        c.File(fullPath)
    })

    // ヘルスチェック
    r.GET("/healthz", func(c *gin.Context){
        c.String(200, "OK")
    })

    log.Println("map-service listening on :8081")
    log.Fatal(r.Run(":8081"))
}
