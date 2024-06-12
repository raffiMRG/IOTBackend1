package controller

import (
	"backendIOT1/Model"
	"backendIOT1/Model/ListMhsModel"
	"backendIOT1/Repository/ListMhsRepository"
	"net/http"

	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
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

// func GetData(c *gin.Context) {
// 	var request ListMahasiswaModel.ListMahasiswa
// 	var response Model.BaseResponseModel

// 	response = ListMahasiswaRepository.GetListMahasiswaByID(request)
// 	if response.CodeResponse != 200 {
// 		c.JSON(response.CodeResponse, response)
// 		return
// 	}

// 	c.JSON(http.StatusOK, response)
// }

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
