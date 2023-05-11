package models

type Pelanggan struct {
	ID              uint   `gorm:"primaryKey;autoIncrement" json:"id_pelanggan"`
	NamaPelanggan   string `json:"nama_pelanggan"`
	JKPelanggan     string `json:"jk_pelanggan"`
	AlamatPelanggan string `json:"alamat_pelanggan"`
	TelpPelanggan   string `json:"telp_pelanggan"`
	FotoPelanggan   string `json:"foto_pelanggan"`
}
