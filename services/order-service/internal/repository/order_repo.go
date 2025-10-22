package repository

import (
	"order-service/internal/model"

	"gorm.io/gorm"
)

// OrderRepository interface for DB ops
type OrderRepository interface {
	CreateOrder(order *model.Order) error
	GetOrderByID(id uint) (*model.Order, error)
	GetOrdersByUser(userID uint, status *model.OrderStatus) ([]model.Order, error)  // Tambah ini
}

// orderRepo struct implements OrderRepository
type orderRepo struct {
	db *gorm.DB
}

// NewOrderRepository creates a new repo instance
func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepo{db: db}
}

func (r *orderRepo) CreateOrder(order *model.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepo) GetOrderByID(id uint) (*model.Order, error) {
	var order model.Order
	err := r.db.Preload("Items").First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *orderRepo) GetOrdersByUser(userID uint, status *model.OrderStatus) ([]model.Order, error) {
	var orders []model.Order
	query := r.db.Preload("Items").Where("user_id = ?", userID)
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	err := query.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}