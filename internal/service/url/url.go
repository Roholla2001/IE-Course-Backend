package url

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	urlmodel "github.com/Roholla2001/ie-course-backend/internal/model/url"
	usermodel "github.com/Roholla2001/ie-course-backend/internal/model/user"
)

type URLServer struct {
	db *gorm.DB
}

func NewURLService(db *gorm.DB) (*URLServer, error) {
	return &URLServer{db}, nil
}

func (us *URLServer) LogRequest(ctx context.Context, id int64, currUser *usermodel.UserModel) error {

	url := new(urlmodel.URLModel)

	if err := us.db.Model(&urlmodel.URLModel{}).Take(url, id).Error; err != nil {
		return err
	}

	if url.UserID != currUser.ID {
		url.FailCount = url.FailCount + 1
		if err := us.db.Model(&urlmodel.URLModel{}).Where("id = ?", url.ID).Save(url).Error; err != nil {
			return err
		}

		return fmt.Errorf("you don't have access to this url")
	}

	url.SuccessCount = url.SuccessCount + 1
	if err := us.db.Model(&urlmodel.URLModel{}).Where("id = ?", url.ID).Save(url).Error; err != nil {
		return err
	}

	return nil
}

func (us *URLServer) GetStats(ctx context.Context, id int64, currUser *usermodel.UserModel) (*urlmodel.URLStat, error) {

	url := new(urlmodel.URLModel)

	if err := us.db.Model(&urlmodel.URLModel{}).Take(url, id).Error; err != nil {
		return nil, err
	}

	if url.UserID != currUser.ID {
		return nil, fmt.Errorf("you don't have access to this url")
	}

	URLStat := new(urlmodel.URLStat)
	if err := us.db.Model(&urlmodel.URLModel{}).Select("success_count", "fail_count").Take(URLStat).Error; err != nil {
		return nil, err
	}

	return URLStat, nil
}

func (us *URLServer) GetRoute(id int64) (string, error) {
	var url string

	if err := us.db.Model(&urlmodel.URLModel{}).Select("url").Take(&url).Error; err != nil {
		return "", err
	}

	return url, nil
}
