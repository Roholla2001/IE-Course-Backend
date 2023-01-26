package adaptor

import (
	"github.com/Roholla2001/ie-course-backend/internal/model"
	"gorm.io/gorm"
)

type ServerController struct {
}

func NewServerController(db *gorm.DB) (*ServerController, error) {
	return &ServerController{}, nil
}

func (sc *ServerController) GetURLs() ([]*model.URL, error){
	return nil, nil
}
