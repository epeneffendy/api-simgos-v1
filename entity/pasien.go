package entity

type Pasien struct {
	Id     int    `json:"id"`
	Norm   string `json:"nomr"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
}
