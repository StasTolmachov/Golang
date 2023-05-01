package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	ID       int     `json:"id"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

var transactions []Transaction
var transactionID int
var transactionMutex sync.Mutex

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/transactions", func(c *gin.Context) {
		c.JSON(http.StatusOK, transactions)
	})

	r.POST("/transactions", func(c *gin.Context) {
		amountStr := c.PostForm("amount")
		currency := c.PostForm("currency")

		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
			return
		}

		transactionMutex.Lock()
		transactionID++
		transaction := Transaction{ID: transactionID, Amount: amount, Currency: currency}
		transactions = append(transactions, transaction)
		transactionMutex.Unlock()

		c.JSON(http.StatusOK, transaction)
	})

	r.Run(":8080")
}
