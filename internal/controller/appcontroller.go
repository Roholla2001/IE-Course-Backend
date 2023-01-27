package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
type AppController struct {
	ServerController *ServerController
	UserController   *UserController
	URLController    []*URLController
}

func (ac *AppController) AddRoutes(parent *gin.RouterGroup) {
	publicRoutes := parent.Group("/")
	privateRoutes := parent.Group("/api")
	privateRoutes.Use(JwtAuthMiddleware())

	ac.UserController.addRoutes(publicRoutes)
	ac.ServerController.addRoutes(privateRoutes)

	for _, c := range ac.URLController {
		c.addRoutes(privateRoutes)
	}
}

func (ac *AppController) NewServerController(db *gorm.DB) (*ServerController, error) {
	return NewServerController(db)
}

func (ac *AppController) NewUserController(db *gorm.DB) (*UserController, error) {
	return NewUserController(db)
}

func (ac *AppController) NewURLController(db *gorm.DB) ([]*URLController, error) {
	if ac.ServerController == nil {
		return nil, fmt.Errorf("can not initiate url controller without a server")
	}

	urls, err := ac.ServerController.getURLs()
	if err != nil {
		return nil, err
	}

	urlControllers := make([]*URLController, 0, len(urls))
	for _, url := range urls {
		uc, err := NewURLController(db, url)
		if err != nil {
			return nil, err
		}

		urlControllers = append(urlControllers, uc)
	}

	return urlControllers, nil
}
