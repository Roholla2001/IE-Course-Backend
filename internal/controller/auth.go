package controller

import (
	"net/http"

	"github.com/Roholla2001/ie-course-backend/internal/infra/datastore"
	usermodel "github.com/Roholla2001/ie-course-backend/internal/model/user"
	userservice "github.com/Roholla2001/ie-course-backend/internal/service/user"
	"github.com/Roholla2001/ie-course-backend/internal/utils/token"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

func CurrentUser(c *gin.Context) (*usermodel.UserModel, error) {

	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		return nil, err
	}

	return userservice.GetUserByID(user_id, datastore.GetDBConn())
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
