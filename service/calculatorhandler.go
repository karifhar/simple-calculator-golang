package service

import (
	"fmt"
	"net/http"
	"web-service-gin/entity"

	"github.com/gin-gonic/gin"
)

// POST Method Calculator
func CalculatorHandler(c *gin.Context) {
	var data entity.Calculator
	var status entity.Status

	if err := c.BindJSON(&data); err != nil {
		fmt.Println("ada error")
		return
	}

	if c.ShouldBind(&data) != nil {
		switch data.Operator {
		case "*":
			status.Result = data.Number1 * data.Number2
		case "+":
			status.Result = data.Number1 + data.Number2
		case "-":
			status.Result = data.Number1 - data.Number2
		case "/":
			if data.Number2 == 0 {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Pembagi tidak boleh Nol"})
				return
			}
			status.Result = data.Number1 / data.Number2
		default:
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "Operator tidak jelas",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "OK",
		"result":  status.Result,
	})
}
