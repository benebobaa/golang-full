package handler

import (
	"middleware_jwt/dto"
	"middleware_jwt/pkg"
	"middleware_jwt/sqlc"
	"strconv"

	"github.com/benebobaa/valo"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	querier sqlc.Querier
}

func NewAuthHandler(q sqlc.Querier) *AuthHandler {
	return &AuthHandler{
		querier: q,
	}
}

func (a *AuthHandler) Register(c *gin.Context) {

	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := valo.Validate(req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if req.Password != req.ConfirmPassword {
		c.JSON(400, gin.H{"error": "password and confirm password must be the same"})
		return
	}

	totalUser, err := a.querier.CountByEmail(c, req.Email)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if totalUser > 0 {
		c.JSON(400, gin.H{"error": "email already registered"})
		return
	}

	hashedPassword, err := pkg.HashPassword(req.Password)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	user, err := a.querier.CreateUser(c, sqlc.CreateUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"data": dto.UserResponse{
		ID:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}})
}

func (a *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := valo.Validate(req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := a.querier.FindByEmail(c, req.Email)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if !pkg.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(400, gin.H{"error": "invalid password"})
		return
	}

	token, err := pkg.GenerateToken(strconv.Itoa(int(user.ID)), user.Email)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": dto.LoginResponse{
		User: dto.UserResponse{
			ID:    int(user.ID),
			Name:  user.Name,
			Email: user.Email,
		},
		Token: token,
	}})
}
