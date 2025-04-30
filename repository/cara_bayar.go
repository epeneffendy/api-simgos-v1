package repository

import (
	"api-simgos/entity"
	"database/sql"
)

type CaraBayarRepository interface {
	GetAllReguler(subSistem int) ([]entity.CaraBayar, error)
	GetAllPartial(subSistem int) ([]entity.CaraBayar, error)
}

type caraBayarRepository struct {
	db *sql.DB
}

func NewCaraBayarRepository(db *sql.DB) *caraBayarRepository {
	return &caraBayarRepository{db}
}

func (c *caraBayarRepository) GetAllReguler(subSistem int) ([]entity.CaraBayar, error) {
	var result []entity.CaraBayar
	sqlStatement := `SELECT kode, nama FROM m_carabayar ORDER BY orders  ASC`
	rows, err := c.db.Query(sqlStatement)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var carabayar entity.CaraBayar
		err := rows.Scan(&carabayar.Id, &carabayar.Nama)
		if err != nil {
			return result, err
		}
		result = append(result, carabayar)
	}
	return result, nil
}

func (c *caraBayarRepository) GetAllPartial(subSistem int) ([]entity.CaraBayar, error) {
	var result []entity.CaraBayar
	sqlStatement := `SELECT kode, nama FROM m_carabayar where kode in (1, 12) order by ORDERS ASC`
	rows, err := c.db.Query(sqlStatement)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var carabayar entity.CaraBayar
		err := rows.Scan(&carabayar.Id, &carabayar.Nama)
		if err != nil {
			return result, err
		}
		result = append(result, carabayar)
	}
	return result, nil
}
