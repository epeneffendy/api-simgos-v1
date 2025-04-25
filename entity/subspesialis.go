package entity

type Subspesialis struct {
	Id              int    `json:"id"`
	Kode            string `json:"kode"`
	Spesialis       string `json:"spesialis"`
	Subspesialis    string `json:"subspesialis"`
	PoliHfis        string `json:"poli_hfis"`
	IdGroupLocation int    `json:"id_group_location"`
}
