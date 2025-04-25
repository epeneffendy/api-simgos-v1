package service

import (
	"api-simgos/entity"
	"api-simgos/repository"
)

type PoliklinikService interface {
	GetBySubSistem(subSistem int) ([]entity.Poliklinik, error)
}

type poliklinikService struct {
	poliklinikRepository repository.PoliklinikRepository
}

func NewPoliklinikService(poliklinikRepository repository.PoliklinikRepository) *poliklinikService {
	return &poliklinikService{poliklinikRepository}
}

func (p *poliklinikService) GetBySubSistem(subSistem int) ([]entity.Poliklinik, error) {
	var subSistemId = subSistem
	var poliklinik []entity.Poliklinik
	var err error
	switch subSistemId {
	case 1, 2, 3:
		poliklinik, err = p.poliklinikRepository.GetBySubSistemReguler(subSistem)
		if err != nil {
			return poliklinik, err
		}
	case 10:
		poliklinik, err = p.poliklinikRepository.GetBySubSistemPenunjang(subSistem)
		if err != nil {
			return poliklinik, err
		}
	case 11:
		poliklinik, err = p.poliklinikRepository.GetBySubSistemTindakan(subSistem)
	default:
		return poliklinik, err
	}
	return poliklinik, nil

}
