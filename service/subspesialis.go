package service

import (
	"api-simgos/entity"
	"api-simgos/repository"
)

type SubSpesialisService interface {
	GetAll() ([]entity.Subspesialis, error)
}

type subSpesialisService struct {
	subspesialisRepository repository.SubspesialisRepository
}

func NewSubSpesialisService(subspesialisRepository repository.SubspesialisRepository) *subSpesialisService {
	return &subSpesialisService{subspesialisRepository}
}

func (p *subSpesialisService) GetAll() ([]entity.Subspesialis, error) {
	subspesialis, err := p.subspesialisRepository.GetAll()
	if err != nil {
		return subspesialis, err
	}
	return subspesialis, nil
}
