package models

type Pembayaran struct {
	ID              uint   `gorm:"primaryKey" json:"id_pembayaran"`
	IDPelanggan     uint   `json:"id_pelanggan"`
	Name            string `json:"name"`
	JenisPembayaran string `json:"jenis_pembayaran"`
	JenisOrder      string `json:"jenis_order"`
}
