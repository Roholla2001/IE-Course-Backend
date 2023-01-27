package server

import (
	"context"
	"fmt"

	urlmodel "github.com/Roholla2001/ie-course-backend/internal/model/url"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

func NewServer(db *gorm.DB) (*Server, error) {
	return &Server{db}, nil
}

func (s *Server) AddURL(ctx context.Context, url *urlmodel.URLModel) error {

	var URLcount int64
	if err := s.DB.Model(&urlmodel.URLModel{}).Select("COUNT(*)").Where("user_id = ?", url.UserID).Take(&URLcount).Error; err != nil {
		return err
	}

	if URLcount >= 20 {
		return fmt.Errorf("url count limit exeeded")
	}

	if err := s.DB.Create(url).Error; err != nil {
		return err
	}

	return nil
}

func (s *Server) GetUserURLs(ctx context.Context, uid int64)([]*string, error){
	urls := make([]*string, 0)

	err := s.DB.Model(&urlmodel.URLModel{}).Where("id = ?", uid).Select("url").Find(&urls).Error
	if err != nil {
		return nil, err
	}

	return urls, nil
}

func (s *Server) GetURLs() ([]*urlmodel.URLModel, error) {
	urls := make([]*urlmodel.URLModel, 0)

	err := s.DB.Model(&urlmodel.URLModel{}).Find(&urls).Error
	if err != nil {
		return nil, err
	}

	return urls, nil
}
