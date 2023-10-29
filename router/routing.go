package router

import (
	"user-personalize/controller"

	"github.com/gin-gonic/gin"
)

func Routing(r *gin.Engine) {
	r.POST("/user/login", controller.Login)
	r.POST("/user/register", controller.Register)
}