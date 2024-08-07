package handler

import (
	"fmt"
	"middleware_jwt/dto"
	"middleware_jwt/middleware"
	"middleware_jwt/pkg"
	"middleware_jwt/sqlc"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	querier sqlc.Querier
}

func NewUserHandler(q sqlc.Querier) *UserHandler {
	return &UserHandler{
		querier: q,
	}
}

func (u *UserHandler) Profile(c *gin.Context) {

	value := c.MustGet(middleware.ClaimsKey).(pkg.UserInfo)

	fmt.Println("value: ", value)

	user, err := u.querier.FindByEmail(c, value.Email)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": dto.UserResponse{
		ID:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}})
}
