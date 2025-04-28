package handler

import (
	"net/http"

	"github.com/atoyr/virtual-arena/map-service/internal/service"
	"github.com/gin-gonic/gin"
)

// mapSvc はマップ操作のビジネスロジックを提供します
var mapSvc = service.NewMapService("./data/maps")

// ListMaps: マップ一覧を返却します
func ListMaps(c *gin.Context) {
	maps, err := mapSvc.ListMaps()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, maps)
}

// GetMap: 指定 ID のマップ JSON を返却します
func GetMap(c *gin.Context) {
	id := c.Param("id")
	m, err := mapSvc.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, m)
}

// CreateMap: 新規マップを作成します（JSON ボディ受け取り）
func CreateMap(c *gin.Context) {
	var req service.Map
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := mapSvc.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// UpdateMap: 既存マップを更新します
func UpdateMap(c *gin.Context) {
	id := c.Param("id")
	var req service.Map
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := mapSvc.Update(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// DeleteMap: 指定 ID のマップを削除します
func DeleteMap(c *gin.Context) {
	id := c.Param("id")
	if err := mapSvc.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
