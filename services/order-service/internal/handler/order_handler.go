package handler

import (
	"net/http"
	"strconv"

	"order-service/internal/model"
	"order-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// OrderHandler sets up routes
type OrderHandler struct {
	service service.OrderService
}

// NewOrderHandler creates handler instance
func NewOrderHandler(svc service.OrderService) *OrderHandler {
	_ = godotenv.Load() // Load .env if needed
	return &OrderHandler{service: svc}
}

// authMiddleware extracts and validates JWT, sets userID to context
func (h *OrderHandler) AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" || len(token) < 7 || token[:7] != "Bearer " {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		c.Abort()
		return
	}
	token = token[7:]

	userID, err := h.service.ValidateJWT(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.Set("userID", userID)  // Simpan ke context untuk digunakan di handler
	c.Set("authHeader", token)
	// c.Set("role", role)
	c.Next()
}

func (h *OrderHandler) CreateOrderHandler(c *gin.Context) {
	userIDAny, exists := c.Get("userID")
	if !exists {
		c.Set("status", http.StatusUnauthorized)
		c.Set("message", "Unauthorized user")
		c.Abort()
		return
	}

	userID := userIDAny.(uint)

	// AMBIL TOKEN DARI CONTEXT
    authHeaderAny, exists := c.Get("authHeader")
    if !exists {
        // Ini seharusnya tidak terjadi jika middleware berjalan
        c.Set("status", http.StatusUnauthorized)
        c.Set("message", "Authorization token not found in context")
        c.Abort()
        return
    }
    authHeader := authHeaderAny.(string)

	var items []service.OrderItemRequest

	if err := c.ShouldBindJSON(&items); err != nil {
		c.Set("status", http.StatusBadRequest)
		c.Set("message", "Invalid request format")
		c.Error(err)
		c.Abort()
		return
	}

	order, err := h.service.CreateOrder(userID, items, authHeader)
	if err != nil {
		c.Set("status", http.StatusInternalServerError)
		c.Set("message", "Failed to create order")
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Order created successfully",
		"data":    order,
	})
}

// GetOrderHandler handles GET
func (h *OrderHandler) GetOrderHandler(c *gin.Context) {
	userIDAny, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID := userIDAny.(uint)

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	order, err := h.service.GetOrder(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	// Check jika order milik user ini (security)
	if order.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

// GetOrdersHandler handles GET /api/orders?status=...
func (h *OrderHandler) GetOrdersHandler(c *gin.Context) {
	userIDAny, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID := userIDAny.(uint)

	var status *model.OrderStatus
	statusStr := c.Query("status")
	if statusStr != "" {
		st := model.OrderStatus(statusStr)
		status = &st
		// Validasi status valid (future: tambah check enum)
	}

	orders, err := h.service.GetOrders(userID, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}