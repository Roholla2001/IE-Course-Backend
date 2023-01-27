package controller

import (
	"net/http"

	usermodel "github.com/Roholla2001/ie-course-backend/internal/model/user"
	userservice "github.com/Roholla2001/ie-course-backend/internal/service/user"
	"github.com/Roholla2001/ie-course-backend/internal/utils/apiutils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	UserServer *userservice.UserServer
}

func NewUserController(db *gorm.DB) (*UserController, error) {
	userService, err := userservice.NewUserService(db)
	if err != nil {
		return nil, err
	}
	return &UserController{UserServer: userService}, nil
}

func (uc *UserController) Register(ctx *gin.Context) {

	var input RegisterInput

	if ok := apiutils.ReadFromJSON(ctx, input); ok {
		return
	}

	u := &usermodel.UserModel{}

	u.Username = input.Username
	u.Password = input.Password

	err := uc.UserServer.CreateUser(ctx, u)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "registration success"})
}

func (uc *UserController) Login(ctx *gin.Context) {
	var input LoginInput

	if ok := apiutils.ReadFromJSON(ctx, input); ok {
		return
	}

	u := &usermodel.UserModel{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := uc.UserServer.LoginCheck(ctx, u)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc *UserController) addRoutes(parent *gin.RouterGroup) {
	parent.POST("/register", uc.Register)
	parent.POST("/login", uc.Login)
}
