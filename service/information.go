package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Information(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Input 2 numbers and 1 operator in postman"})
}
