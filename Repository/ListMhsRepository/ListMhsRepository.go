package ListMhsRepository

import (
	"backendIOT1/Model"
	connection "backendIOT1/Model/Connection"
	"backendIOT1/Model/ListMhsModel"
	"fmt"
)

func InsertListMahasiswa(ListMahasiswa ListMhsModel.ListMhs) Model.BaseResponseModel {
	var query string
	var result Model.BaseResponseModel
	db := connection.DB
	query = "INSERT INTO tb_card (uuid_card, nama_mhs) VALUES (?, ?)"

	tempResult := db.Exec(query, ListMahasiswa.UUID_CARD, ListMahasiswa.Nama_mhs)

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

func UpdateListMahasiswa(ListMahasiswa ListMhsModel.ListMhs) Model.BaseResponseModel {
	var query string
	var result Model.BaseResponseModel
	db := connection.DB
	query = "UPDATE tb_card SET nama_mhs = ? WHERE uuid_card = ?"

	tempResult := db.Exec(query, ListMahasiswa.Nama_mhs, ListMahasiswa.UUID_CARD)

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
	query = "DELETE FROM tb_card WHERE uuid_card = ?"

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
		query = "SELECT * FROM tb_card WHERE uuid_card = ?"
		tempResult = db.Where("uuid_card = ?", request.UUID_CARD).Find(&ListMahasiswaList)
		fmt.Println("uuid_card query")
	} else {
		query = "SELECT * FROM tb_card"
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
