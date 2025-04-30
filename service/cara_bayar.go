package service

import (
	"api-simgos/entity"
	"api-simgos/repository"
)

type CaraBayarService interface {
	GetAll(subSistem int) ([]entity.CaraBayar, error)
}

type caraBayarService struct {
	caraBayarRepository repository.CaraBayarRepository
}

func NewCaraBayarService(caraBayarRepository repository.CaraBayarRepository) *caraBayarService {
	return &caraBayarService{caraBayarRepository}
}

func (c *caraBayarService) GetAll(subSistem int) ([]entity.CaraBayar, error) {
	var subSistemId = subSistem
	var caraBayar []entity.CaraBayar
	var err error

	switch subSistemId {
	case 1, 2, 3:
		caraBayar, err = c.caraBayarRepository.GetAllReguler(subSistem)
		if err != nil {
			return caraBayar, nil
		}
	case 10, 11:
		caraBayar, err = c.caraBayarRepository.GetAllPartial(subSistem)
		if err != nil {
			return caraBayar, nil
		}
	default:
		return caraBayar, nil
	}

	return caraBayar, nil
}
