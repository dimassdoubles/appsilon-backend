package main

import (
	"net/http"

	"appsilon.id/mdtrns/transaction"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"title": "api payment gateway appsilon"})
	})

	r.POST("/transaction", transaction.CreateTransaction)

	r.Run()
}
