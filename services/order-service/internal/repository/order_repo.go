package repository

import (
	"order-service/internal/model"

	"gorm.io/gorm"
)

// OrderRepository interface for DB ops
type OrderRepository interface {
	CreateOrder(order *model.Order) error
	GetOrderByID(id uint) (*model.Order, error)
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