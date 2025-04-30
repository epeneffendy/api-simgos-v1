package repository

import (
	"api-simgos/entity"
	"database/sql"
)

type SubspesialisRepository interface {
	GetAll() ([]entity.Subspesialis, error)
}
type subspesialisRepository struct {
	db2 *sql.DB
}

func NewSubspesialisRepository(db2 *sql.DB) *subspesialisRepository {
	return &subspesialisRepository{db2}
}

func (p *subspesialisRepository) GetAll() ([]entity.Subspesialis, error) {
	var result []entity.Subspesialis
	sqlStatement := `select sp.id, ss."Kode" as kode ,  ss."Spesialis" as spesialis,sp."Kode" as poli_hfis, sp."Subspesialis" as subspesialis, ss."Id_group_location_simgos" as id_group_location 
					from daftar_spesialis ss  
					join daftar_subspesialis sp on ss.id = sp."FK_daftar_spesialis_ID" 
					group by ss."Kode" ,  ss."Spesialis", sp.id, sp."Subspesialis", sp."Kode", ss."Id_group_location_simgos"`
	rows, err := p.db2.Query(sqlStatement)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var subspesialis entity.Subspesialis
		var idGroup sql.NullString
		err := rows.Scan(&subspesialis.Id, &subspesialis.Kode, &subspesialis.Spesialis, &subspesialis.PoliHfis, &subspesialis.Subspesialis, &idGroup)

		if err == sql.ErrNoRows {
			return result, err
		}

		if err != nil {
			return result, err
		}
		result = append(result, subspesialis)
	}
	return result, nil
}
