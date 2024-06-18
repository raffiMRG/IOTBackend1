package AbsenMhsModel

import "time"

type ListAbsen struct {
	// ID         int    `json:"id" form:"id"`
	UUID_CARD string    `json:"uuid_card" form:"uuid_card"`
	Nama_mhs  string    `json:"nama_mhs" form:"nama_mhs"`
	CreateAt  time.Time `gorm:"column:create_at;autoCreateTime"` // autoCreateTime untuk otomatis mengisi waktu pembuatan
	// CREATE_AT string `json:"create_at" form:"create_at"`
	// NIM        string `json:"nim" form:"nim"`
}

// TableName memberikan nama tabel yang eksplisit
func (ListAbsen) TableName() string {
	return "tb_absen"
}
