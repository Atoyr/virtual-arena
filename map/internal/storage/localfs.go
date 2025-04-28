package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type LocalFS struct {
	BasePath string
}

func NewLocalFS(basePath string) *LocalFS {
	return &LocalFS{BasePath: basePath}
}

const (
	mapPath = "static/maps"
)

func (l *LocalFS) SaveMapJSON(tenantID, mapID string, data []byte) error {
	dir := filepath.Join(l.BasePath, tenantID, "maps", mapID)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, "map.json"), data, 0644)
}

// LoadMap loads the map.json file for the given mapID
func (l *LocalFS) LoadMap(mapID string) ([]byte, error) {
	return os.ReadFile(filepath.Join(l.BasePath, mapPath, mapID, "map.json"))
}

func (l *LocalFS) SaveTile(tenantID, mapID, z, x, y string, data io.Reader) error {
	dir := filepath.Join(l.BasePath, tenantID, "maps", mapID, "tiles", z, x)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(dir, y+".png"))
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, data)
	return err
}

func (l *LocalFS) LoadTileset(mapID, tilesetID string) ([]byte, error) {
	filename := fmt.Sprintf("%s.png", tilesetID)
	return os.ReadFile(filepath.Join(l.BasePath, mapPath, mapID, "tilesets", filename))
}
