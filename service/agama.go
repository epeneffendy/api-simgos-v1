package service

import (
	"api-simgos/entity"
	"api-simgos/repository"
)

type AgamaService interface {
	GetAll() ([]entity.Agama, error)
}

type agamaService struct {
	agamaRepository repository.AgamaRepository
}

func NewAgamaService(agamaRepository repository.AgamaRepository) *agamaService {
	return &agamaService{agamaRepository}
}

func (c *agamaService) GetAll() ([]entity.Agama, error) {
	agama, err := c.agamaRepository.GetAll()
	if err != nil {
		return agama, err
	}
	return agama, nil
}
