// internal/storage/storage.go
package storage

import "io"

type Storage interface {
  // map.json 用
  SaveMapJSON(tenantID, mapID string, data []byte) error
  LoadMapJSON(tenantID, mapID string) ([]byte, error)

  // タイル画像用
	LoadTileset(tenantID, tilesetId string) ([]byte, error)
}
