package model

type Map struct {
	ID       string  `json:"id"`
	Width    int     `json:"width"`
	Height   int     `json:"height"`
	TileSize int     `json:"tileSize"`
	Layers   []Layer `json:"layers"`
}

type Layer struct {
	Name string `json:"name"`
	Data []int  `json:"data"`
}
