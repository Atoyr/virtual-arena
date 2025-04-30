package service

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/atoyr/virtual-arena/map-service/internal/model"
)

type MapService struct {
	basePath string
}

func NewMapService(basePath string) *MapService {
	return &MapService{basePath: basePath}
}

// ListMaps returns a list of maps
func (s *MapService) ListMaps() ([]model.Map, error) {
	// basePath 配下のフォルダを列挙してマップ一覧を構築
	dirs, err := os.ReadDir(s.basePath)
	if err != nil {
		return nil, err
	}
	var out []model.Map
	for _, d := range dirs {
		if d.IsDir() {
			raw, _ := os.ReadFile(filepath.Join(s.basePath, d.Name(), "map.json"))
			var m model.Map
			json.Unmarshal(raw, &m)
			out = append(out, m)
		}
	}
	return out, nil
}

// GetMap returns a map by ID
func (s *MapService) GetMap(id string) (*model.Map, error) {
	path := filepath.Join(s.basePath, id, "map.json")
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var m model.Map
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// CreateMap creates a new map
func (s *MapService) CreateMap(m *model.Map) (*model.Map, error) {
	if m.ID == "" {
		m.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	}
	dir := filepath.Join(s.basePath, m.ID)
	os.MkdirAll(dir, 0755)
	// m.Updated = time.Now()
	raw, _ := json.MarshalIndent(m, "", "  ")
	os.WriteFile(filepath.Join(dir, "map.json"), raw, 0644)
	return m, nil
}

// UpdateMpa updates an existing map
func (s *MapService) UpdateMap(id string, m *model.Map) (*model.Map, error) {
	dir := filepath.Join(s.basePath, id)
	m.ID = id
	// m.Updated = time.Now()
	raw, _ := json.MarshalIndent(m, "", "  ")
	os.WriteFile(filepath.Join(dir, "map.json"), raw, 0644)
	return m, nil
}

// DeleteMap deletes a map by ID
func (s *MapService) DeleteMap(id string) error {
	return os.RemoveAll(filepath.Join(s.basePath, id))
}
