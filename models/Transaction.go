package models

type Transaction struct {
	ID                uint    `gorm:"primaryKey" json:"id_transaksi"`
	IDPelanggan       uint    `json:"id_pelanggan"`
	UserID            uint    `json:"user_id"`
	IDJenis           uint    `json:"id_jenis"`
	TglTerima         string  `json:"tgl_terima"`
	TglSelesai        string  `json:"tgl_selesai"`
	JmlKilo           float64 `json:"jml_kilo"`
	AlamatPengantaran string  `json:"alamat_pengantaran"`
	Catatan           string  `json:"catatan"`
	TotalBayar        float64 `json:"totalbayar"`
	StatusPembayaran  string  `json:"status_pembayaran"`
	StatusPengambilan string  `json:"status_pengambilan"`
}
