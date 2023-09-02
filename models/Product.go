package models

type Product struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	NamaProduk string `gorm:"type:varchar(50)" json:"nama_produk"`
	Deskripsi  string `gorm:"type:varchar(50)" json:"deskripsi"`
}
