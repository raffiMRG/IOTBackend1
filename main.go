package main

import (
	controller "backendIOT1/Controller"
	connection "backendIOT1/Model/Connection"

	"github.com/gin-gonic/gin"
)

func main() {
	port := ":1211"
	r := gin.Default()
	connection.ConnectDatabase()

	//###BEGIN WEB API
	// Get data
	r.GET("/api/ListMahasiswa/GetData", controller.GetData)

	// Insert data
	r.POST("/api/ListMahasiswa/InsertData", controller.InsertData)

	// Update data
	r.PUT("/api/ListMahasiswa/UpdateData", controller.UpdateData)

	// Delete data
	r.DELETE("/api/ListMahasiswa/DeleteData", controller.DeleteData)
	//###END WEB API

	r.Run(port)
}
