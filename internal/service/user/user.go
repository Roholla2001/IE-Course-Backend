package user

import (
	"context"
	"fmt"

	usermodel "github.com/Roholla2001/ie-course-backend/internal/model/user"
	auth "github.com/Roholla2001/ie-course-backend/internal/utils/token"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type UserServer struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) (*UserServer, error) {
	return &UserServer{db}, nil
}

func (us *UserServer) CreateUser(ctx context.Context, user *usermodel.UserModel) error {
	if err := us.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (us *UserServer) LoginCheck(ctx context.Context, user *usermodel.UserModel) (token string, err error) {
	var u usermodel.UserModel

	//take user with given credintials
	if err = us.db.Model(&usermodel.UserModel{}).Where("userـname = ?", user.UserName).Take(&u).Error; err != nil {
		return
	}

	err = usermodel.VerifyPassword(user.Password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return
	}
	if err != nil {
		return
	}

	token, err = auth.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func GetUserByID(uid int64, db *gorm.DB) (*usermodel.UserModel, error) {
	var u *usermodel.UserModel

	if err := db.First(u, uid).Error; err != nil {
		return u, fmt.Errorf("user not found")
	}

	u.Password = ""

	return u, nil

}
