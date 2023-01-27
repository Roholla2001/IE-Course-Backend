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
	server      *server.Server
	parentRoute *gin.RouterGroup
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

	url := new(urlmodel.URLModel)
	if ok := apiutils.ReadFromJSON(ctx, url); !ok {
		return
	}

	currUser, err := CurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url.UserID = currUser.ID

	err = sc.server.AddURL(c, url)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uc, err := NewURLController(sc.server.DB, url)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sc.parentRoute.GET("/"+url.URL, uc.Log)
	sc.parentRoute.GET("/"+url.URL+"/stats", uc.GetStats)

	ctx.JSON(http.StatusOK, gin.H{"message": "url added successfully"})

}

func (sc *ServerController) GetUserURLs(ctx *gin.Context){
	c := ctx.Request.Context()

	currUser, err := CurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	urls, err := sc.server.GetUserURLs(c, currUser.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	apiutils.WriteToJSON(ctx, urls, err)
}

func (sc *ServerController) getURLs() ([]*urlmodel.URLModel, error) {
	return sc.server.GetURLs()
}

func (sc *ServerController) addRoutes(parent *gin.RouterGroup) {
	sc.parentRoute = parent
	parent.POST("/add-url", sc.AddUrl)
	parent.GET("/get-urls", sc.GetUserURLs)
}
