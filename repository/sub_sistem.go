package repository

import (
	"api-simgos/entity"
	"database/sql"
)

type SubSistemRepository interface {
	GetAll() ([]entity.SubSistem, error)
}

type subSistemRepository struct {
	db *sql.DB
}

func NewSubSistemRepository(db *sql.DB) *subSistemRepository {
	return &subSistemRepository{db}
}

func (p *subSistemRepository) GetAll() ([]entity.SubSistem, error) {
	var result []entity.SubSistem
	sqlStatement := `SELECT id_subsistem as id, nama_subsistem FROM m_subsistem where (sistem='IRJA' or sistem ='R.JALAN' or sistem='R.DARURAT') order by id_subsistem ASC `
	rows, err := p.db.Query(sqlStatement)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var subsistem entity.SubSistem
		err := rows.Scan(&subsistem.Id, &subsistem.NamaSubSistem)
		if err != nil {
			return result, err
		}
		result = append(result, subsistem)
	}
	return result, nil
}
