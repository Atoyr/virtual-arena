package storage

import (
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "path/filepath"
)

type LocalFS struct {
    BasePath string
}

func NewLocalFS(basePath string) *LocalFS {
    return &LocalFS{BasePath: basePath}
}

func (l *LocalFS) SaveMapJSON(tenantID, mapID string, data []byte) error {
    dir := filepath.Join(l.BasePath, tenantID, "maps", mapID)
    if err := os.MkdirAll(dir, 0755); err != nil {
        return err
    }
    return ioutil.WriteFile(filepath.Join(dir, "map.json"), data, 0644)
}

func (l *LocalFS) LoadMapJSON(tenantID, mapID string) ([]byte, error) {
    return ioutil.ReadFile(filepath.Join(l.BasePath, tenantID, "maps", mapID, "map.json"))
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

func (l *LocalFS) TileURL(tenantID, mapID, z, x, y string) (string, error) {
    // 開発時はローカルサーバから直接参照
    return fmt.Sprintf("/maps/%s/tiles/%s/%s/%s.png", mapID, z, x, y), nil
}

