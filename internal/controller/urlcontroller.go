package controller

import (
	"net/http"

	urlmodel "github.com/Roholla2001/ie-course-backend/internal/model/url"
	urlservice "github.com/Roholla2001/ie-course-backend/internal/service/url"
	"github.com/Roholla2001/ie-course-backend/internal/utils/apiutils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type URLController struct {
	URLID     int64
	URLServer *urlservice.URLServer
}

func NewURLController(db *gorm.DB, url *urlmodel.URL) (*URLController, error) {
	us, err := urlservice.NewURLService(db)
	if err != nil {
		return nil, err
	}
	return &URLController{URLID: url.ID, URLServer: us}, nil
}

func (cc *URLController) Log(ctx *gin.Context) {
	c := ctx.Request.Context()

	var url *urlmodel.URL
	if ok := apiutils.ReadFromJSON(ctx, url); !ok {
		return
	}

	currUser, err := CurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = cc.URLServer.LogRequest(c, currUser)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (cc *URLController) GetStats(ctx *gin.Context) {
	c := ctx.Request.Context()

	var url *urlmodel.URL
	if ok := apiutils.ReadFromJSON(ctx, url); !ok {
		return
	}

	currUser, err := CurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stat ,err := cc.URLServer.GetStats(c, currUser)
	if err != nil {
		return
	}

	apiutils.WriteToJSON(ctx, stat, err)
}

func (cc *URLController) getRoute() string {
	url, _ := cc.URLServer.GetRoute(cc.URLID)
	return url
}

func (cc *URLController) addRoutes(parent *gin.RouterGroup) {
	//url endpoints
	ep := parent

	url := cc.getRoute()
	ep.Use()
	{
		ep.GET("/"+url, cc.Log)
		ep.GET("/"+url+"/stats", cc.GetStats)
	}
}
