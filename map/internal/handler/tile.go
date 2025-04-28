package handler

import (
	"net/http"

	"github.com/atoyr/virtual-arena/map-service/internal/service"
	"github.com/atoyr/virtual-arena/map-service/internal/storage"
	"github.com/gin-gonic/gin"
)

// mapSvc はマップ操作のビジネスロジックを提供します
var tileSvc = service.NewTileService(storage.NewLocalFS(""))

// ListTiles: マップに紐付くタイルセット一覧を返却します
func ListTilesets(c *gin.Context) {
	maps, err := mapSvc.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, maps)
}
