package transaction

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

const BASE_URL = "https://app.sandbox.midtrans.com/snap/v1/transactions"

const SERVER_KEY = "{SERVER_KEY}"

type CreateTransactionInput struct {
	TransactionDetails transactionDetails `json:"transaction_details"`
}

type transactionDetails struct {
	OrderId     string `json:"order_id"`
	GrossAmount int    `json:"gross_amount"`
}

func CreateTransaction(context *gin.Context) {
	var input CreateTransactionInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jsonBody, err := json.Marshal(input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create request
	req, err := http.NewRequest("POST", BASE_URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// set the request header
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// add authorization header
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(SERVER_KEY+":")))

	// send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer resp.Body.Close()

	// read response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": string(respBody)})
}
