package controller

import (
	urlmodel "github.com/Roholla2001/ie-course-backend/internal/model/url"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type URLController struct {
}

func NewURLController(db *gorm.DB, url *urlmodel.URL) (*URLController, error) {
	return &URLController{}, nil
}

func (c *URLController) Log(ctx *gin.Context) {

}

func (c *URLController) GetStats(ctx *gin.Context) {

}

func (c *URLController) getRoute() string {
	return ""
}

func (c *URLController) addRoutes(parent *gin.RouterGroup) {
	//url endpoints
	ep := parent.Group("/link")

	url := c.getRoute()
	ep.Use()
	{
		ep.GET("/"+url, c.Log)
		ep.GET("/"+url+"/stats", c.GetStats)
	}
}
