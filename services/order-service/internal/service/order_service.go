package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"order-service/internal/model"
	"order-service/internal/repository"

	"github.com/golang-jwt/jwt/v5"
)

// OrderService interface
type OrderService interface {
	CreateOrder(userID uint, items []OrderItemRequest, authHeader string) (*model.Order, error)
	GetOrder(id uint) (*model.Order, error)
	GetOrders(userID uint, status *model.OrderStatus) ([]model.Order, error)  // Tambah ini
	ValidateJWT(tokenString string) (uint, error)
}

// OrderItemRequest for incoming payload
type OrderItemRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

// orderService struct
type orderService struct {
	repo       repository.OrderRepository
	productURL string
	userURL    string
	jwtSecret  string
}

// NewOrderService creates service instance
func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{
		repo:       repo,
		productURL: os.Getenv("PRODUCT_SERVICE_URL"),
		userURL:    os.Getenv("USER_SERVICE_URL"),
		jwtSecret:  os.Getenv("JWT_SECRET"),
	}
}

// CreateOrder validates and creates order, links to Product
func (s *orderService) CreateOrder(userID uint, items []OrderItemRequest, authHeader string) (*model.Order, error) {
	var totalPrice float64
	orderItems := make([]model.OrderItem, len(items))

	for i, item := range items {
		// Call Product Service to get product details and check stock
		product, err := s.getProduct(item.ProductID)
		if err != nil {
			return nil, err
		}
		if product.Stock < item.Quantity {
			return nil, errors.New("insufficient stock for product: " + item.ProductID)
		}

		price := product.Price * float64(item.Quantity)
		totalPrice += price
		orderItems[i] = model.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		}
	}

	order := &model.Order{
		UserID:     userID,
		TotalPrice: totalPrice,
		Items:      orderItems,
	}

	if err := s.repo.CreateOrder(order); err != nil {
		return nil, err
	}

	log.Println("Order created in DB, attempting to update stock...")

	for _, item := range items {
		if err := s.UpdateProductStock(item.ProductID, -item.Quantity, authHeader); err != nil {
            log.Printf("Failed to update stock... %v", err)
            // PENTING: Kembalikan error-nya agar transaksi gagal
            return nil, err
        }
    }

	// Future: Async update stock in Product via message broker
	log.Println("Stock updated successfully.")
	return order, nil
}

// UpdateProductStock calls Product Service to patch stock
func (s *orderService) UpdateProductStock(productID string, delta int, authHeader string) error {
	url := fmt.Sprintf("%s/api/products/%s", s.productURL, productID)
	body := map[string]int{"stock": delta}  // Negative untuk reduce
	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + authHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Periksa jika status code BUKAN di rentang 2xx (sukses)
    if resp.StatusCode < 200 || resp.StatusCode >= 300 {
        // Coba baca body error jika ada, untuk log yang lebih baik
        // (Ini opsional tapi sangat membantu)
		log.Printf("Failed to update product stock, status: %d", resp.StatusCode)
        return errors.New("failed to update product stock")
    }

	return nil
}

// GetOrder fetches order details
func (s *orderService) GetOrder(id uint) (*model.Order, error) {
	return s.repo.GetOrderByID(id)
}

func (s *orderService) GetOrders(userID uint, status *model.OrderStatus) ([]model.Order, error) {
	return s.repo.GetOrdersByUser(userID, status)
}

// getProduct calls Product Service API (reusable HTTP client)
func (s *orderService) getProduct(id string) (*ProductResponse, error) {
	url := fmt.Sprintf("%s/api/products/%s", s.productURL, id)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to call Product Service: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("product not found")
	}

	var wrapper struct {
		Product ProductResponse `json:"product"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Product response: %+v", wrapper.Product)
	return &wrapper.Product, nil
}

// ProductResponse from Product Service
type ProductResponse struct {
	ID    string  `json:"_id"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

func (s *orderService) ValidateJWT(tokenString string) (uint, error) {
	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}
	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims in token")
	}

	if claims["iss"] != "paring-user-service" {
		return 0, errors.New("invalid issuer")
	}

	// Get email (sub) from claims
	email, ok := claims["sub"].(string)
	if !ok || email == "" {
		return 0, errors.New("missing email (sub) in token claims")
	}

	log.Printf("[ValidateJWT] Token validated for email: %s", email)

	exp, ok := claims["exp"].(float64)
	if !ok || time.Unix(int64(exp), 0).Before(time.Now()) {
		return 0, errors.New("token expired")
	}

	userIdFloat, ok := claims["userId"].(float64)
	if !ok {
		return 0, errors.New("missing userId in token")
	}

	role, ok := claims["role"].(string)
	if !ok || (role != "USER" && role != "SELLER") {  // Adjust roles di Paring
		return 0, errors.New("invalid role")
	}

	// Prepare request to User Service
	url := fmt.Sprintf("%s/api/users/profile", s.userURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+tokenString)

	// Execute HTTP request
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[ValidateJWT] Failed to call User Service: %v", err)
		return 0, fmt.Errorf("failed to connect to user service: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("user profile not found (status: %d)", resp.StatusCode)
	}

	// Decode JSON response
	var userResp UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&userResp); err != nil {
		return 0, fmt.Errorf("failed to parse user service response: %v", err)
	}

	log.Printf("[ValidateJWT] User ID resolved: %d for email: %s", userResp.ID, email)
	// return uint(userResp.ID), nil
	return uint(userIdFloat), nil
}

// UserResponse from User Service
type UserResponse struct {
	ID uint `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Role string `json:"role"`
	// Tambah fields lain jika needed
}