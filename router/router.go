package router

import (
	"project2/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/people", controllers.GetPeople)
	r.POST("/people", controllers.InsertPerson)
	r.GET("/people/:id", controllers.GetPersonByID)
	r.POST("/people/:id", controllers.DeleteUserbyID)

	//-------------------------------------------------//
	r.GET("/fruit", controllers.GetFruits)
	r.POST("/fruit", controllers.InsertFruit)

}
