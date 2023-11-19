package routers

import (
	"calculator_backend/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	router := r.Group("")
	{
		router.PUT("/register", controllers.Register)
		router.GET("/login", controllers.Login)
		router.GET("/get_histories", controllers.GetHistories)
		router.PUT("/add_history", controllers.AddHistory)
		router.GET("/get_rate", controllers.GetRateRecord)
		router.PUT("/add_rate", controllers.AddRateRecord)

	}
}
