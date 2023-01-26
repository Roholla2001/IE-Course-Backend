package adaptor

import "gorm.io/gorm"

type UserController struct {
}

func NewUserController(db *gorm.DB) (*UserController,error){
	return &UserController{}, nil
}

