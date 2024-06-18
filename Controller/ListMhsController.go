package controller

import (
	"backendIOT1/Model"
	"backendIOT1/Model/AbsenMhsModel"
	"backendIOT1/Model/ListMhsModel"
	"backendIOT1/Repository/ListMhsRepository"
	"net/http"

	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetData(c *gin.Context) {
	var requestData ListMhsModel.ListMhs
	var response Model.BaseResponseModel

	// Read the body of the request
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read request body"})
		return
	}

	// Unmarshal the JSON data
	json.Unmarshal(body, &requestData)
	// if err := json.Unmarshal(body, &requestData); err != nil {
	// c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
	// return
	// }

	response = ListMhsRepository.GetListMahasiswaByID(requestData)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	// Use the parsed data
	fmt.Printf("Received data: %+v\n", requestData)

	// Respond with the received data
	c.JSON(http.StatusOK, response)
	// c.JSON(http.StatusOK, gin.H{
	// 	"id":   requestData.ID,
	// 	"nim":  requestData.NIM,
	// 	"nama": requestData.Nama,
	// })
}

func GetAbsen(c *gin.Context) {
	var requestAbsen AbsenMhsModel.ListAbsen
	var response Model.BaseResponseModel

	// Read the body of the request
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read request body"})
		return
	}

	// Unmarshal the JSON data
	json.Unmarshal(body, &requestAbsen)

	response = ListMhsRepository.GetAbsenMahasiswaByID(requestAbsen)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	// Use the parsed data
	fmt.Printf("Received data: %+v\n", requestAbsen)

	// Respond with the received data
	c.JSON(http.StatusOK, response)
}

func InsertData(c *gin.Context) {
	var request ListMhsModel.ListMhs
	var response Model.BaseResponseModel

	if err := c.ShouldBindJSON(&request); err != nil {
		response = Model.BaseResponseModel{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = ListMhsRepository.InsertListMahasiswa(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

type AbsenController struct {
	repository ListMhsRepository.AbsenRepository
}

func NewAbsenController(db *gorm.DB) *AbsenController {
	repo := ListMhsRepository.NewAbsenRepository(db)
	return &AbsenController{repository: repo}
}

func (ctrl *AbsenController) CopyNama(c *gin.Context) {
	var input struct {
		UUID_CARD string `json:"uuid_card"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	akun, err := ctrl.repository.FindByUUIDCard(input.UUID_CARD)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	absen := AbsenMhsModel.ListAbsen{
		UUID_CARD: akun.UUID_CARD,
		Nama_mhs:  akun.Nama_mhs,
	}
	if err := ctrl.repository.CreateAbsen(&absen); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Record inserted successfully"})
}

// func InsertAbsen(c *gin.Context) {
// if err := c.ShouldBindJSON(&request); err != nil {
// 	response = Model.BaseResponseModel{
// 		CodeResponse:  400,
// 		HeaderMessage: "Bad Request",
// 		Message:       err.Error(),
// 		Data:          nil,
// 	}
// 	c.JSON(http.StatusBadRequest, response)
// 	return
// }

// SUCCESS SCENARIO
// response = ListMhsRepository.InsertAbsenMahasiswa(request, source)
// if response.CodeResponse != 200 {
// 	c.JSON(response.CodeResponse, response)
// 	return
// }

// c.JSON(http.StatusOK, response)
// }

func UpdateData(c *gin.Context) {
	var request ListMhsModel.ListMhs
	var response Model.BaseResponseModel

	if err := c.ShouldBindJSON(&request); err != nil {
		response = Model.BaseResponseModel{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = ListMhsRepository.UpdateListMahasiswa(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func DeleteData(c *gin.Context) {
	var request ListMhsModel.ListMhs
	var response Model.BaseResponseModel

	if err := c.ShouldBindJSON(&request); err != nil {
		response = Model.BaseResponseModel{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = ListMhsRepository.DeleteListMahasiswa(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
