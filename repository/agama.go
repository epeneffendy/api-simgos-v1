package repository

import (
	"api-simgos/entity"
	"database/sql"
)

type AgamaRepository interface {
	GetAll() ([]entity.Agama, error)
}

type agamaRepository struct {
	db *sql.DB
}

func NewAgamaRepository(db *sql.DB) *agamaRepository {
	return &agamaRepository{db}
}

func (a *agamaRepository) GetAll() ([]entity.Agama, error) {
	var result []entity.Agama
	sqlStatement := `SELECT * from m_agama_pasien order by id_agama ASC`
	rows, err := a.db.Query(sqlStatement)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var agama entity.Agama
		err := rows.Scan(&agama.Id, &agama.Nama)
		if err != nil {
			return result, err
		}

		result = append(result, agama)

	}
	return result, nil
}
