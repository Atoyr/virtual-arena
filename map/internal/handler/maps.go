package handler

import (
	"net/http"

	"github.com/atoyr/virtual-arena/map-service/internal/service"
	"github.com/gin-gonic/gin"
)

func GetMaps(svc *service.MapService) gin.HandlerFunc {
	return func(c *gin.Context) {
		maps, err := svc.ListMaps()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, maps)
	}
}

func GetMapMeta(svc *service.MapService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("mapID")
		m, err := svc.Get(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "map not found"})
			return
		}
		c.JSON(http.StatusOK, m)
	}
}
