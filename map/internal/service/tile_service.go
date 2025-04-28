package service

import (
	"github.com/atoyr/virtual-arena/map-service/internal/storage"
)

type Tileset struct {
	GID           int    `json:"firstgid"`
	Name          string `json:"name"`
	TileWidth     int    `json:"width"`
	TileHeight    int    `json:"height"`
	Margin        int    `json:"margin"`
	Spacing       int    `json:"spacing"`
	Columns       int    `json:"columns"`
	TileCount     int    `json:"tilecount"`
	ImageFileName string `json:"image"`
	ImageWidth    int    `json:"imagewidth"`
	ImageHeight   int    `json:"imageheight"`
	Tiles         []Tile `json:"tiles"`
}

type Tile struct {
	ID         int        `json:"id"`
	Properties []Property `json:"properties"`
}

type TileService struct {
	storage storage.Storage
}

func NewTileService(storage storage.Storage) *TileService {
	return &TileService{storage: storage}
}

func (s *TileService) Tilesets() ([]Tileset, error) {
	// FIXME: 実装

	t := Tileset{
		GID:           1,
		Name:          "test",
		TileWidth:     32,
		TileHeight:    32,
		Margin:        0,
		Spacing:       0,
		Columns:       8,
		TileCount:     192,
		ImageFileName: "test.png",
		ImageWidth:    256,
		ImageHeight:   768,
		Tiles:         []Tile{{ID: 0, Properties: []Property{{Name: "test", Type: "string", Value: "test"}}}},
	}
	ts := []Tileset{t}
	return ts, nil

}
