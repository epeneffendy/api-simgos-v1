package repository

import (
	"api-simgos/entity"
	"database/sql"
)

type PasienRepository interface {
	GetAll() ([]entity.Pasien, error)
}

type pasienRepository struct {
	db *sql.DB
}

func NewPasienRepository(db *sql.DB) *pasienRepository {
	return &pasienRepository{db}
}

func (p *pasienRepository) GetAll() ([]entity.Pasien, error) {
	var result []entity.Pasien
	sqlStatement := `SELECT id, nomr, nama, alamat FROM m_pasien limit 10`
	rows, err := p.db.Query(sqlStatement)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var pasien entity.Pasien
		err := rows.Scan(&pasien.Id, &pasien.Norm, &pasien.Nama, &pasien.Alamat)
		if err != nil {
			return result, err
		}
		result = append(result, pasien)
	}

	return result, nil
}
