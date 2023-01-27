package controller

import (
	"net/http"

	urlmodel "github.com/Roholla2001/ie-course-backend/internal/model/url"
	"github.com/Roholla2001/ie-course-backend/internal/service/server"
	"github.com/Roholla2001/ie-course-backend/internal/utils/apiutils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServerController struct {
	server *server.Server
}

func NewServerController(db *gorm.DB) (*ServerController, error) {
	s, err := server.NewServer(db)
	if err != nil {
		return nil, err
	}
	return &ServerController{server: s}, nil
}

func (sc *ServerController) AddUrl(ctx *gin.Context) {
	c := ctx.Request.Context()

	url := new(urlmodel.URL)
	if ok := apiutils.ReadFromJSON(ctx, url); !ok {
		return
	}

	err := sc.server.AddURL(c, url)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (sc *ServerController) getURLs() ([]*urlmodel.URL, error) {
	return nil, nil
}

func (sc *ServerController) addRoutes(parent *gin.RouterGroup) {
	parent.POST("/add-url", sc.AddUrl)
}
