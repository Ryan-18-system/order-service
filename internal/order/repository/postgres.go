package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/Ryan-18-system/order-service/internal/order/domain"
)

type postgresOrderRepository struct {
	db *gorm.DB
}

func NewPostgresOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &postgresOrderRepository{db: db}
}

func (r *postgresOrderRepository) List(ctx context.Context) ([]domain.Order, error) {
	var orders []domain.Order
	if err := r.db.WithContext(ctx).Table("orders").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
