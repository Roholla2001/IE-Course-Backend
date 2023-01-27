package router

import (
	"github.com/Roholla2001/ie-course-backend/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(ac *controller.AppController) *gin.Engine {
	router := gin.Default()
	ac.AddRoutes(&router.RouterGroup)

	return router
}
