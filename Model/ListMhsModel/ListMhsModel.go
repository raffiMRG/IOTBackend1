package ListMhsModel

type ListMhs struct {
	// KTPID      string `json:"ktpId" form:"ktpId"`
	// NamaTamu   string `json:"namaTamu" form:"namaTamu"`
	// Tujuan     string `json:"tujuan" form:"tujuan"`
	// Keterangan string `json:"keterangan" form:"keterangan"`
	ID         int    `json:"id" form:"id"`
	UUID_CARD  string `json:"uuid_card" form:"uuid_card"`
	Nama_mhs   string `json:"nama_mhs" form:"nama_mhs"`
	CREATED_AT string `json:"created_at" form:"created_at"`
}

// TableName memberikan nama tabel yang eksplisit
func (ListMhs) TableName() string {
	return "tb_card"
}
