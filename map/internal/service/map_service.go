package service

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Map struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Width   int       `json:"width"`
	Height  int       `json:"height"`
	Tileset string    `json:"tileset"`
	Updated time.Time `json:"updated"`
}

type MapService struct {
	basePath string
}

func NewMapService(basePath string) *MapService {
	return &MapService{basePath: basePath}
}

func (s *MapService) List() ([]Map, error) {
	// basePath 配下のフォルダを列挙してマップ一覧を構築
	dirs, err := os.ReadDir(s.basePath)
	if err != nil {
		return nil, err
	}
	var out []Map
	for _, d := range dirs {
		if d.IsDir() {
			raw, _ := os.ReadFile(filepath.Join(s.basePath, d.Name(), "map.json"))
			var m Map
			json.Unmarshal(raw, &m)
			out = append(out, m)
		}
	}
	return out, nil
}

func (s *MapService) Get(id string) (*Map, error) {
	path := filepath.Join(s.basePath, id, "map.json")
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var m Map
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *MapService) Create(m *Map) (*Map, error) {
	if m.ID == "" {
		m.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	}
	dir := filepath.Join(s.basePath, m.ID)
	os.MkdirAll(dir, 0755)
	m.Updated = time.Now()
	raw, _ := json.MarshalIndent(m, "", "  ")
	os.WriteFile(filepath.Join(dir, "map.json"), raw, 0644)
	return m, nil
}

func (s *MapService) Update(id string, m *Map) (*Map, error) {
	dir := filepath.Join(s.basePath, id)
	m.ID = id
	m.Updated = time.Now()
	raw, _ := json.MarshalIndent(m, "", "  ")
	os.WriteFile(filepath.Join(dir, "map.json"), raw, 0644)
	return m, nil
}

func (s *MapService) Delete(id string) error {
	return os.RemoveAll(filepath.Join(s.basePath, id))
}
