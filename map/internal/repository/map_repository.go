package repository

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/atoyr/virtual-arena/map-service/internal/model"
)

// MapRepository はマップ情報の永続化アクセスインターフェースを定義します
type MapRepository interface {
	// FindAll はすべてのマップを返します
	FindAll() ([]model.Map, error)
	// FindByID は指定 ID のマップを返します。見つからない場合はエラーを返します
	FindByID(id string) (*model.Map, error)
}

// fileMapRepository はファイルベースの MapRepository 実装です
type fileMapRepository struct {
	dir string
}

// NewMapRepository はディレクトリパスを指定してファイル実装を生成します
func NewMapRepository(dir string) MapRepository {
	return &fileMapRepository{dir: dir}
}

// FindAll はディレクトリ内の JSON ファイルをすべて読み込み、Map スライスとして返します
func (r *fileMapRepository) FindAll() ([]model.Map, error) {
	files, err := os.ReadDir(r.dir)
	if err != nil {
		return nil, err
	}
	maps := make([]model.Map, 0, len(files))
	for _, f := range files {
		if f.IsDir() || filepath.Ext(f.Name()) != ".json" {
			continue
		}
		path := filepath.Join(r.dir, f.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		var m model.Map
		if err := json.Unmarshal(data, &m); err != nil {
			return nil, err
		}
		maps = append(maps, m)
	}
	return maps, nil
}

// FindByID は指定された ID の JSON ファイルを読み込み、Map を返します
func (r *fileMapRepository) FindByID(id string) (*model.Map, error) {
	filePath := filepath.Join(r.dir, id+".json")
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("map not found")
		}
		return nil, err
	}
	var m model.Map
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	return &m, nil
}
