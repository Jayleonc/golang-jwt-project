package routes

import (
	controller "github.com/Jayleonc/golang-jwt-project/controllers"
	middleware "github.com/Jayleonc/golang-jwt-project/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
