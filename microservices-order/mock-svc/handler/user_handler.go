package handler

import "github.com/gin-gonic/gin"

type User struct {
	ID            string `json:"id"`
	AccountBankID string `json:"account_bank_id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
}

type UserHandler struct {
	db map[string]User
}

func (h *UserHandler) GetUser(c *gin.Context) {
	var users []User

	for _, user := range h.db {
		users = append(users, user)
	}

	c.JSON(200, gin.H{"data": users})
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, ok := h.db[id]
	if !ok {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	c.JSON(200, gin.H{"data": user})
}
