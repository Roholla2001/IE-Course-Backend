package router

import (
	"fmt"

	"github.com/Roholla2001/ie-course-backend/internal/adaptor"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppController struct {
	ServerController *adaptor.ServerController
	UserController   *adaptor.UserController
	URLController    []*adaptor.URLController
}

func InitRouter(ac *AppController) *gin.Engine {
	router := gin.Default()
	ac.AddRoutes(&router.RouterGroup)

	return router
}

func (ac *AppController) AddRoutes(parent *gin.RouterGroup) {
}

func (ac *AppController) NewServerController(db *gorm.DB) (*adaptor.ServerController, error) {
	return adaptor.NewServerController(db)
}

func (ac *AppController) NewUserController(db *gorm.DB) (*adaptor.UserController, error) {
	return adaptor.NewUserController(db)
}


func (ac *AppController) NewURLController(db *gorm.DB) ([]*adaptor.URLController, error) {
	if ac.ServerController == nil {
		return nil, fmt.Errorf("Cannot initiate url controller without a server")
	}

	urls, err := ac.ServerController.GetURLs()
	if err != nil {
		return nil, err
	}

	urlControllers := make([]*adaptor.URLController, 0, len(urls))
	for _, url := range urls {
		uc, err := adaptor.NewURLController(db, url)
		if err != nil {
			return nil, err
		}

		urlControllers = append(urlControllers, uc)
	}

	return urlControllers, nil
}
