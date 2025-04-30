package handler

import (
	"net/http"

	"github.com/atoyr/virtual-arena/map-service/internal/controller"
	"github.com/gin-gonic/gin"
)

func GetMapsHandler(mapController *controller.MapController) gin.HandlerFunc {
	return func(c *gin.Context) {
		maps, err := mapController.ListMaps()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, maps)
	}
}

func GetMapMetaHandler(mapController *controller.MapController) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("mapID")
		m, err := mapController.GetMap(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "map not found"})
			return
		}
		c.JSON(http.StatusOK, m)
	}
}
