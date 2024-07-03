package ListMhsModel

import "time"

type ListMhs struct {
	UUID_CARD string    `json:"uuid_card" form:"uuid_card"`
	NIM       string    `json:"nim" form:"nim"`
	Nama_mhs  string    `json:"nama_mhs" form:"nama_mhs"`
	CREATE_AT time.Time `gorm:"column:create_at;autoCreateTime"` // autoCreateTime untuk otomatis mengisi waktu pembuatan
	// CREATE_AT string `json:"create_at" form:"create_at"`
}

// TableName memberikan nama tabel yang eksplisit
func (ListMhs) TableName() string {
	return "tb_akun"
}
