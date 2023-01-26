package adaptor

import (
	"github.com/Roholla2001/ie-course-backend/internal/model"
	"gorm.io/gorm"
)

type URLController struct {
}

func NewURLController(db *gorm.DB, url *model.URL) (*URLController, error)