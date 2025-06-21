package routes

import (
	"github.com/JagdeepSingh13/go_jwt/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.Engine) {
	route.POST("users/signup", controllers.Signup())
	route.POST("users/login", controllers.Login())
}
