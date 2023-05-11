package models

import (
	"time"
)

type Laundry struct {
	ID                uint      `gorm:"primaryKey" json:"id_laundry"`
	IDPelanggan       uint      `json:"id_pelanggan"`
	UserID            uint      `json:"user_id"`
	IDJenis           uint      `json:"id_jenis"`
	TglTerima         time.Time `json:"tgl_terima"`
	TglSelesai        time.Time `json:"tgl_selesai"`
	JmlKilo           float64   `json:"jml_kilo"`
	AlamatPengantaran string    `json:"alamat_pengantaran"`
	Catatan           string    `json:"catatan"`
	TotalBayar        float64   `json:"totalbayar"`
	StatusPembayaran  string    `json:"status_pembayaran"`
	StatusPengambilan string    `json:"status_pengambilan"`
}
