package service

import (
	"api-simgos/entity"
	"api-simgos/repository"
)

type SubSistemService interface {
	GetAll() ([]entity.SubSistem, error)
}

type subSistemService struct {
	subSistemRepository repository.SubSistemRepository
}

func NewSubSistemService(subSistemRepository repository.SubSistemRepository) *subSistemService {
	return &subSistemService{subSistemRepository}
}

func (p *subSistemService) GetAll() ([]entity.SubSistem, error) {
	subSistem, err := p.subSistemRepository.GetAll()
	if err != nil {
		return subSistem, err
	}
	return subSistem, nil
}
