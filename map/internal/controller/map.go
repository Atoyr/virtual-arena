package controller

import (
	"github.com/atoyr/virtual-arena/map-service/internal/model"
	"github.com/atoyr/virtual-arena/map-service/internal/service"
)

func NewController(svc *service.MapService) *MapController {
	return &MapController{
		service: svc,
	}
}

type MapController struct {
	service *service.MapService
}

func (mc *MapController) ListMaps() ([]model.Map, error) {
	return mc.service.ListMaps()
}

func (mc *MapController) GetMap(id string) (*model.Map, error) {
	return mc.service.GetMap(id)
}
