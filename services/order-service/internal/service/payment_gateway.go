package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

// PaymentGateway interface (segregated untuk confirm saja)
type PaymentGateway interface {
	ConfirmPayment(orderID uint, amount float64) (string, error)  // Return transactionID jika success
}

// mockPaymentGateway implements mock
type mockPaymentGateway struct {
	url string
}

// NewMockPaymentGateway creates instance
func NewMockPaymentGateway() PaymentGateway {
	return &mockPaymentGateway{url: os.Getenv("PAYMENT_MOCK_URL")}
}

func (m *mockPaymentGateway) ConfirmPayment(orderID uint, amount float64) (string, error) {
	url := fmt.Sprintf("%s/api/payment/confirm", m.url)
	body := map[string]interface{}{
		"orderId": orderID,
		"amount":  amount,
	}
	jsonBody, _ := json.Marshal(body)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Printf("Failed to call payment mock: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if result["status"] != "SUCCESS" {
		return "", errors.New(result["message"])
	}

	return result["transactionId"], nil
}