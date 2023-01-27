package adaptor

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