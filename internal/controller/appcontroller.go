package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
type AppController struct {
	ServerController *ServerController
	UserController   *UserController
	URLController    []*URLController
}

type testUrl struct {
	Url string `json:"url"`
}

func (ac *AppController) AddRoutes(parent *gin.RouterGroup) {
	route := parent.Group("/api")

	route.POST("/add-route", func(ctx *gin.Context) {
		// c := ctx.Request.Context()

		req, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.Error(err)
		}

		test := new(testUrl)
		err = json.Unmarshal(req, test)
		if err != nil {
			ctx.Error(err)
		}

		route.GET("/"+test.Url, func(ctx *gin.Context) {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "url works!"})
		})

		if err != nil {
			ctx.Error(err)
		}

		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "url added", "url": test.Url})
	})
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

	urls, err := ac.ServerController.GetURLs()
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
