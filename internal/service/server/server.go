package server

import (
	"context"
	"fmt"

	urlmodel "github.com/Roholla2001/ie-course-backend/internal/model/url"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
}

func NewServer(db *gorm.DB) (*Server, error) {
	return &Server{db}, nil
}

func (s *Server) AddURL(ctx context.Context, url *urlmodel.URL) error {

	var URLcount int64
	if err := s.db.Model(&urlmodel.URL{}).Select("COUNT(*)").Where("user_id = ?", url.UserID).Take(&URLcount).Error; err != nil {
		return err
	}

	if URLcount >= 20 {
		return fmt.Errorf("url count limit exeeded")
	}

	if err := s.db.Create(url).Error; err != nil {
		return err
	}

	return nil
}

func (s *Server) GetURLs() ([]*urlmodel.URL, error) {
	urls := make([]*urlmodel.URL, 0)

	err := s.db.Model(&urlmodel.URL{}).Find(urls).Error
	if err != nil {
		return nil, err
	}

	return urls, nil
}
