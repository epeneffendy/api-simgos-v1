package repository

import (
	"api-simgos/entity"
	"database/sql"
)

type PoliklinikRepository interface {
	GetBySubSistemReguler(subsistem int) ([]entity.Poliklinik, error)
	GetBySubSistemPenunjang(subsistem int) ([]entity.Poliklinik, error)
	GetBySubSistemTindakan(subsistem int) ([]entity.Poliklinik, error)
}

type poliklinikRepository struct {
	db *sql.DB
}

func NewPoliklinikRepository(db *sql.DB) *poliklinikRepository {
	return &poliklinikRepository{db}
}

func (p *poliklinikRepository) GetBySubSistemReguler(subsistem int) ([]entity.Poliklinik, error) {
	var result []entity.Poliklinik
	sqlStatement := `select a.no as kode, a.nama from m_ruang a
					inner join m_subsistem b on b.nama_subsistem = a.subsistem
					where b.id_subsistem = $1 `
	rows, err := p.db.Query(sqlStatement, subsistem)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var poliklinik entity.Poliklinik
		err := rows.Scan(&poliklinik.Kode, &poliklinik.Nama)
		if err != nil {
			return result, err
		}
		result = append(result, poliklinik)
	}

	return result, nil
}

func (p *poliklinikRepository) GetBySubSistemPenunjang(subsistem int) ([]entity.Poliklinik, error) {
	var result []entity.Poliklinik
	sqlStatement := `select a.no as kode, a.nama from m_ruang a
					inner join m_subsistem b on b.nama_subsistem = a.is_partial 
					where b.id_subsistem = $1`
	rows, err := p.db.Query(sqlStatement, subsistem)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var poliklinik entity.Poliklinik
		err := rows.Scan(&poliklinik.Kode, &poliklinik.Nama)
		if err != nil {
			return result, err
		}
		result = append(result, poliklinik)
	}

	return result, nil
}

func (p *poliklinikRepository) GetBySubSistemTindakan(subsistem int) ([]entity.Poliklinik, error) {
	var result []entity.Poliklinik
	sqlStatement := `SELECT form_id as kode,nama_field as nama from m_form_lab where  unit_form = '70' and jenis_kelompok = 'Ruang Tindakan'  and aktif = 1 order by nama_field asc`
	rows, err := p.db.Query(sqlStatement, subsistem)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var poliklinik entity.Poliklinik
		err := rows.Scan(&poliklinik.Kode, &poliklinik.Nama)
		if err != nil {
			return result, err
		}
		result = append(result, poliklinik)
	}

	return result, nil
}
