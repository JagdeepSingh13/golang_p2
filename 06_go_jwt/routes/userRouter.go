package routes

import (
	"github.com/JagdeepSingh13/go_jwt/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.RouterGroup) {
	// route.Use(middleware.Authenticate())
	route.GET("/users", controllers.GetUsers())
	route.GET("users/:user_id", controllers.GetUser())
}
