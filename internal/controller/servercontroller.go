package controller

import (
	urlmodel "github.com/Roholla2001/ie-course-backend/internal/model/url"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServerController struct {
}

func NewServerController(db *gorm.DB) (*ServerController, error) {
	return &ServerController{}, nil
}

func (sc *ServerController) AddUrl(ctx *gin.Context)

func (sc *ServerController) getURLs() ([]*urlmodel.URL, error) {
	return nil, nil
}

func (sc *ServerController) addRoutes(parent *gin.RouterGroup) {
	sr := parent.Group("/")

	sr.POST("/add-url", sc.AddUrl)
}
