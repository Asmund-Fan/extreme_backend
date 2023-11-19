package main

import (
	"calculator_backend/db"
	"calculator_backend/routers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db.InitDb()

	router := gin.Default()
	routers.InitRouter(router)

	err := router.Run(":2333")
	if err != nil {
		log.Printf("Having problem starting gin.Engine: %v \n", err)
		return
	}
}
