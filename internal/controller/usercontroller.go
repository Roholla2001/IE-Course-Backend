package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
}

func NewUserController(db *gorm.DB) (*UserController,error){
	return &UserController{}, nil
}

func Register(ctx *gin.Context){

}

func Login(ctx *gin.Context){
	
}