package router

import (
	"net/http"

	"itv-go/internal/auth"

	"github.com/gin-gonic/gin"
)

func fakeLoginHandler(c *gin.Context) {
	fakeUsername := "hello"

	token, err := auth.GenerateJWT(fakeUsername)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
