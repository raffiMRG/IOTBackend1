package main

import (
	controller "backendIOT1/Controller"
	connection "backendIOT1/Model/Connection"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	port := ":1211"
	r := gin.Default()
	connection.ConnectDatabase()
	db, err := connection.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	// Create controller
	absenController := controller.NewAbsenController(db)

	//###BEGIN WEB API
	// Get data
	r.GET("/api/ListMahasiswa/GetData", controller.GetData)
	r.GET("/api/ListAbsen/GetData", controller.GetAbsen)

	// Insert data
	r.POST("/api/ListMahasiswa/InsertData", controller.InsertData)
	r.POST("/api/ListAbsen/InsertData", absenController.CopyNama)

	// Update data
	r.PUT("/api/ListMahasiswa/UpdateData", controller.UpdateData)

	// Delete data
	r.DELETE("/api/ListMahasiswa/DeleteData", controller.DeleteData)
	//###END WEB API

	r.Run(port)
}
