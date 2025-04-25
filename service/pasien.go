package service

import (
	"api-simgos/entity"
	"api-simgos/repository"
)

type PasienService interface {
	GetAll() ([]entity.Pasien, error)
}

type pasienService struct {
	pasienRepository repository.PasienRepository
}

func NewPasienService(pasienRepository repository.PasienRepository) *pasienService {
	return &pasienService{pasienRepository}
}

func (p *pasienService) GetAll() ([]entity.Pasien, error) {
	pasien, err := p.pasienRepository.GetAll()
	if err != nil {
		return pasien, err
	}
	return pasien, nil
}
