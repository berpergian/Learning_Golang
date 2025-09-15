package staticdata

import (
	"log"

	"github.com/berpergian/chi_learning/shared/constant"
	"github.com/berpergian/chi_learning/shared/model"
)

type CharacterStaticDataRepo struct{ svc *StaticDataService }
type ItemStaticDataRepo struct{ svc *StaticDataService }

func (s *StaticDataService) Characters() CharacterStaticDataRepo { return CharacterStaticDataRepo{s} }
func (s *StaticDataService) Items() ItemStaticDataRepo           { return ItemStaticDataRepo{s} }

func (r CharacterStaticDataRepo) Get(id string) (model.CharacterStaticData, bool) {
	val, ok := r.svc.Get(constant.CharacterStaticData, id)
	if !ok {
		return model.CharacterStaticData{}, false
	}
	ch, ok := val.(*model.CharacterStaticData)
	if !ok {
		return model.CharacterStaticData{}, false
	}
	return *ch, true
}

func (r CharacterStaticDataRepo) All() []model.CharacterStaticData {
	vals := r.svc.GetAll(constant.CharacterStaticData)
	log.Printf("[Static Data] Character All: %d", len(vals))

	res := make([]model.CharacterStaticData, 0, len(vals))
	for _, v := range vals {
		if ch, ok := v.(*model.CharacterStaticData); ok {
			res = append(res, *ch) // deref to return value
		} else {
			log.Printf("[Static Data] unexpected type in Character All: %T", v)
		}
	}
	return res
}

func (r ItemStaticDataRepo) Get(id string) (model.ItemStaticData, bool) {
	val, ok := r.svc.Get(constant.ItemStaticData, id)
	if !ok {
		return model.ItemStaticData{}, false
	}
	ch, ok := val.(*model.ItemStaticData)
	if !ok {
		return model.ItemStaticData{}, false
	}
	return *ch, true
}

func (r ItemStaticDataRepo) All() []model.ItemStaticData {
	vals := r.svc.GetAll(constant.ItemStaticData)
	log.Printf("[Static Data] Item All: %d", len(vals))

	res := make([]model.ItemStaticData, 0, len(vals))
	for _, v := range vals {
		if ch, ok := v.(*model.ItemStaticData); ok {
			res = append(res, *ch)
		} else {
			log.Printf("[Static Data] unexpected type in Item All: %T", v)
		}
	}
	return res
}
