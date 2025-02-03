package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func HandleLogin(c *gin.Context) {
	var loginReq LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Add real authentication logic here
	// This is just a mock response for testing
	if loginReq.Username == "test" && loginReq.Password == "password" {
		token := "mock-jwt-token" // In reality, generate a proper JWT token
		c.JSON(http.StatusOK, User{
			Username: loginReq.Username,
			Token:    token,
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}
