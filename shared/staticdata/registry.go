package staticdata

import (
	"sync"

	"github.com/berpergian/chi_learning/shared/model"
)

type Factory func() model.StaticModel

type StaticDataService struct {
	data     map[string]map[string]model.StaticModel
	registry map[string]Factory // folder -> factory
	mu       sync.RWMutex
}

func InitializeStaticData() *StaticDataService {
	staticDataService := NewStaticDataService()

	staticDataService.Register("character", func() model.StaticModel { return &model.CharacterStaticData{} })
	staticDataService.Register("item", func() model.StaticModel { return &model.ItemStaticData{} })

	return staticDataService
}

func NewStaticDataService() *StaticDataService {
	return &StaticDataService{
		data:     make(map[string]map[string]model.StaticModel),
		registry: make(map[string]Factory),
	}
}

// Register a new folder with its struct type factory
func (s *StaticDataService) Register(folder string, factory Factory) {
	s.registry[folder] = factory
	s.data[folder] = make(map[string]model.StaticModel)
}
