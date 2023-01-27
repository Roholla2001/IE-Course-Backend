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

func (uc *UserController) Register(ctx *gin.Context){

}

func (uc *UserController) Login(ctx *gin.Context){

}

func (uc *UserController) addRoutes(parent *gin.RouterGroup){
	
}