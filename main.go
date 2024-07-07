package main

import (
	"log"

	"project2/database"
	"project2/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Veritabanı bağlantısını başlat
	database.InitDB()
	defer database.CloseDB()

	// Gin router oluştur
	r := gin.Default()

	// Router'ı yapılandır
	router.SetupRouter(r)

	log.Println("Server is running on port 8080")
	log.Fatal(r.Run(":8080"))
}
