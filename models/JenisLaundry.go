package models

type JenisLaundry struct {
	ID           uint    `gorm:"primaryKey" json:"id_jenis"`
	JenisLaundry string  `json:"jenis_laundry"`
	LamaProses   string  `json:"lama_proses"`
	Tarif        float64 `json:"tarif"`
}
