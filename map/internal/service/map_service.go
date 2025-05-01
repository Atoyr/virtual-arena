package service

import (
	"errors"

	"github.com/atoyr/virtual-arena/map-service/internal/model"
	"github.com/atoyr/virtual-arena/map-service/internal/repository"
)

type MapService struct {
	repo repository.MapRepository
}

func NewMapService(repo repository.MapRepository) *MapService {
	return &MapService{repo: repo}
}

// ListMaps returns a list of maps
func (s *MapService) ListMaps() ([]model.Map, error) {
	maps, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return maps, nil
}

// GetMap returns a map by ID
func (s *MapService) GetMapByID(id string) (*model.Map, error) {
	m, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if m == nil {
		return nil, errors.New("map not found")
	}
	return m, nil
}

// // CreateMap creates a new map
// func (s *MapService) CreateMap(m *model.Map) (*model.Map, error) {
// 	if m.ID == "" {
// 		m.ID = fmt.Sprintf("%d", time.Now().UnixNano())
// 	}
// 	dir := filepath.Join(s.basePath, m.ID)
// 	os.MkdirAll(dir, 0755)
// 	// m.Updated = time.Now()
// 	raw, _ := json.MarshalIndent(m, "", "  ")
// 	os.WriteFile(filepath.Join(dir, "map.json"), raw, 0644)
// 	return m, nil
// }
//
// // UpdateMpa updates an existing map
// func (s *MapService) UpdateMap(id string, m *model.Map) (*model.Map, error) {
// 	dir := filepath.Join(s.basePath, id)
// 	m.ID = id
// 	// m.Updated = time.Now()
// 	raw, _ := json.MarshalIndent(m, "", "  ")
// 	os.WriteFile(filepath.Join(dir, "map.json"), raw, 0644)
// 	return m, nil
// }
//
// // DeleteMap deletes a map by ID
// func (s *MapService) DeleteMap(id string) error {
// 	return os.RemoveAll(filepath.Join(s.basePath, id))
// }
