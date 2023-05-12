package models

import (
	"time"
)

type Laporan struct {
	ID         uint      `gorm:"primaryKey" json:"id_laporan"`
	TglLaporan time.Time `json:"tgl_laporan"`
	KetLaporan string    `json:"ket_laporan"`
	Catatan    string    `json:"catatan"`
	IDLaundry  uint      `json:"id_laundry"`
	Pemasukan  float64   `json:"pemasukan"`
}
