package staticdata

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/berpergian/chi_learning/shared/model"
)

func (s *StaticDataService) Load(basePath string) error {
	return filepath.WalkDir(basePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}

		folder := filepath.Base(filepath.Dir(path)) // e.g. "character", "item"
		factory, ok := s.registry[folder]
		if !ok {
			// skip unknown folders
			return nil
		}

		file, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		obj := factory()
		if err := json.Unmarshal(file, obj); err != nil {
			return err
		}

		s.store(folder, obj)
		return nil
	})
}

func (s *StaticDataService) store(folder string, obj model.StaticModel) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[folder][obj.GetId()] = obj
}

func (s *StaticDataService) Get(folder, id string) (model.StaticModel, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[folder][id]
	return val, ok
}

func (s *StaticDataService) GetAll(folder string) []model.StaticModel {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var result []model.StaticModel
	for _, v := range s.data[folder] {
		result = append(result, v)
	}
	return result
}
