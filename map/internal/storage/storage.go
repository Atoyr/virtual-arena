// internal/storage/storage.go
package storage

type Storage interface {
	// map.json 用
	LoadMap(mapID string) ([]byte, error)
	// ListMaps() ([]string, error)

	// タイル画像用
	LoadTileset(mapID, tilesetID string) ([]byte, error)
	// ListTilesets() ([]string, error)
}
