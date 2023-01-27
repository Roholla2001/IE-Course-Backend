package controller

import (
	urlmodel "github.com/Roholla2001/ie-course-backend/internal/model/url"
	"gorm.io/gorm"
)

type ServerController struct {
}

func NewServerController(db *gorm.DB) (*ServerController, error) {
	return &ServerController{}, nil
}

func (sc *ServerController) GetURLs() ([]*urlmodel.URL, error) {
	return nil, nil
}
