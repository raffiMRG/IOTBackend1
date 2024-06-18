package ListMhsRepository

import (
	"backendIOT1/Model"
	"backendIOT1/Model/AbsenMhsModel"
	connection "backendIOT1/Model/Connection"
	"backendIOT1/Model/ListMhsModel"
	"fmt"

	"gorm.io/gorm"
)

func InsertListMahasiswa(ListMahasiswa ListMhsModel.ListMhs) Model.BaseResponseModel {
	var query string
	var result Model.BaseResponseModel
	db := connection.DB
	query = "INSERT INTO tb_akun (uuid_card) VALUES (?)"

	tempResult := db.Exec(query, ListMahasiswa.UUID_CARD)

	if tempResult.Error != nil {
		result = Model.BaseResponseModel{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		result = Model.BaseResponseModel{
			CodeResponse:  200,
			HeaderMessage: "Success",
			Message:       "Data berhasil ditambahkan.",
			Data:          ListMahasiswa.UUID_CARD,
		}
	}

	return result
}

// INSERT DATA TO TABLE ABSEN
type AbsenRepository interface {
	FindByUUIDCard(uuidCard string) (*ListMhsModel.ListMhs, error)
	CreateAbsen(absen *AbsenMhsModel.ListAbsen) error
}

type absenRepository struct {
	db *gorm.DB
}

func NewAbsenRepository(db *gorm.DB) AbsenRepository {
	return &absenRepository{db}
}

func (r *absenRepository) FindByUUIDCard(uuidCard string) (*ListMhsModel.ListMhs, error) {
	var akun ListMhsModel.ListMhs
	if err := r.db.Where("uuid_card = ?", uuidCard).First(&akun).Error; err != nil {
		return nil, err
	}
	return &akun, nil
}

// func (r *absenRepository) CreateAbsen(absen *AbsenMhsModel.ListAbsen) error {
// 	if err := r.db.Omit("CreateAt").Create(absen).Error; err != nil { // Gunakan Omit untuk mengabaikan kolom CreateAt
// 		return err
// 	}
// 	return nil
// }

func (r *absenRepository) CreateAbsen(absen *AbsenMhsModel.ListAbsen) error {
	if err := r.db.Create(absen).Error; err != nil {
		return err
	}
	return nil
}

// func InsertAbsenMahasiswa() {

// }

func UpdateListMahasiswa(ListMahasiswa ListMhsModel.ListMhs) Model.BaseResponseModel {
	var query string
	var result Model.BaseResponseModel
	db := connection.DB
	query = "UPDATE tb_akun SET nama_mhs = ? , nim = ? WHERE uuid_card = ?"

	tempResult := db.Exec(query, ListMahasiswa.Nama_mhs, ListMahasiswa.NIM, ListMahasiswa.UUID_CARD)

	if tempResult.Error != nil {
		result = Model.BaseResponseModel{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		rowsAffected := tempResult.RowsAffected
		if rowsAffected == 0 {
			result = Model.BaseResponseModel{
				CodeResponse:  404,
				HeaderMessage: "Not Found",
				Message:       "Data dengan uuid_card tersebut tidak ditemukan.",
				Data:          nil,
			}
		} else {
			result = Model.BaseResponseModel{
				CodeResponse:  200,
				HeaderMessage: "Success",
				Message:       "Data berhasil diubah.",
				Data:          ListMahasiswa.UUID_CARD,
			}
		}
	}

	return result
}

func DeleteListMahasiswa(request ListMhsModel.ListMhs) Model.BaseResponseModel {
	var query string
	var result Model.BaseResponseModel
	db := connection.DB
	query = "DELETE FROM tb_akun WHERE uuid_card = ?"

	tempResult := db.Exec(query, request.UUID_CARD)

	if tempResult.Error != nil {
		result = Model.BaseResponseModel{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		// Periksa apakah ada baris yang terpengaruh oleh perintah DELETE
		rowsAffected := tempResult.RowsAffected
		if rowsAffected == 0 {
			result = Model.BaseResponseModel{
				CodeResponse:  404,
				HeaderMessage: "Not Found",
				Message:       "Data dengan uuid_card tersebut tidak ditemukan.",
				Data:          nil,
			}
		} else {
			result = Model.BaseResponseModel{
				CodeResponse:  200,
				HeaderMessage: "Success",
				Message:       "Data berhasil dihapus.",
				Data:          request.UUID_CARD,
			}
		}
	}

	return result
}

func GetListMahasiswaByID(request ListMhsModel.ListMhs) Model.BaseResponseModel {
	var query string
	var result Model.BaseResponseModel
	var ListMahasiswaList []ListMhsModel.ListMhs
	db := connection.DB
	tempResult := connection.DB

	fmt.Printf("Received data repository: %+v\n", request)

	if request.UUID_CARD != "" {
		query = "SELECT * FROM tb_akun WHERE uuid_card = ?"
		tempResult = db.Where("uuid_card = ?", request.UUID_CARD).Find(&ListMahasiswaList)
		fmt.Println("uuid_card query")
	} else {
		query = "SELECT * FROM tb_akun"
		tempResult = db.Raw(query).Find(&ListMahasiswaList)
		fmt.Println(tempResult)
	}

	// tempResult = db.Raw(query).Find(&ListMahasiswaList)
	// tempResult := db.Exec(query, request.NIM)

	if tempResult.Error != nil {
		result = Model.BaseResponseModel{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		result = Model.BaseResponseModel{
			CodeResponse:  200,
			HeaderMessage: "Success",
			Message:       "Data retrieved successfully",
			Data:          ListMahasiswaList,
		}
	}

	return result
}

func GetAbsenMahasiswaByID(request AbsenMhsModel.ListAbsen) Model.BaseResponseModel {
	var query string
	var result Model.BaseResponseModel
	var ListAbsenList []AbsenMhsModel.ListAbsen
	db := connection.DB
	tempResult := connection.DB

	fmt.Printf("Received data repository: %+v\n", request)

	if request.UUID_CARD != "" {
		query = "SELECT * FROM tb_absen WHERE uuid_card = ?"
		tempResult = db.Where("uuid_card = ?", request.UUID_CARD).Find(&ListAbsenList)
		fmt.Println("uuid_card query")
	} else {
		query = "SELECT * FROM tb_absen"
		tempResult = db.Raw(query).Find(&ListAbsenList)
		fmt.Println(tempResult)
	}

	if tempResult.Error != nil {
		result = Model.BaseResponseModel{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		result = Model.BaseResponseModel{
			CodeResponse:  200,
			HeaderMessage: "Success",
			Message:       "Data retrieved successfully",
			Data:          ListAbsenList,
		}
	}

	return result
}
