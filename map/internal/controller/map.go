package controller

import (
	"github.com/atoyr/virtual-arena/map-service/internal/model"
	"github.com/atoyr/virtual-arena/map-service/internal/service"
)

type MapController struct {
	service *service.MapService
}

func (mc *MapController) ListMaps() ([]model.Map, error) {
	return mc.service.ListMaps()
}
