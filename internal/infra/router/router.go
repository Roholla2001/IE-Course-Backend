package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppController struct {
	User    
}

func InitRouter(ac *AppController) *gin.Engine {
	router := gin.Default()
	ac.AddRoutes(&router.RouterGroup)

	return router
}

func (ac *AppController) AddRoutes(parent *gin.RouterGroup) {
}


func NewUserController(db *gorm.DB)

func NewURLController(db *gorm.DB)

