package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"order-service/internal/model"
	"order-service/internal/repository"

	"github.com/golang-jwt/jwt/v5"
)

// OrderService interface
// OrderService interface
type OrderService interface {
	CreateOrder(userID uint, items []OrderItemRequest) (*model.Order, error)
	GetOrder(id uint) (*model.Order, error)
	ValidateJWT(tokenString string) (uint, error)  // Tambah ini
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
	jwtSecret  string
}

// NewOrderService creates service instance
func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{
		repo:       repo,
		productURL: os.Getenv("PRODUCT_SERVICE_URL"),
		jwtSecret:  os.Getenv("JWT_SECRET"),
	}
}

// CreateOrder validates and creates order, links to Product
func (s *orderService) CreateOrder(userID uint, items []OrderItemRequest) (*model.Order, error) {
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

	// Future: Async update stock in Product via message broker
	return order, nil
}

// GetOrder fetches order details
func (s *orderService) GetOrder(id uint) (*model.Order, error) {
	return s.repo.GetOrderByID(id)
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

// ValidateJWT extracts userID from token (reusable auth)
func (s *orderService) ValidateJWT(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	// claims, ok := token.Claims.(jwt.MapClaims)
	// if !ok {
	// 	return 0, errors.New("invalid claims")
	// }

	// email := claims["sub"].(string) // Asumsi User Service JWT punya sub=email; adjust jika needed
	// Future: Call User Service to get userID by email
	// Untuk simple, asumsi userID dari claim atau hardcode test
	return 1, nil // Ganti dengan real logic
}