package usecase

import (
	"context"

	"github.com/Ryan-18-system/order-service/internal/order/domain"
)

type OrderUseCase interface {
	ListOrders(ctx context.Context) ([]domain.Order, error)
}

type orderUseCase struct {
	repo domain.OrderRepository
}

func NewOrderUseCase(r domain.OrderRepository) OrderUseCase {
	return &orderUseCase{repo: r}
}

func (uc *orderUseCase) ListOrders(ctx context.Context) ([]domain.Order, error) {
	return uc.repo.List(ctx)
}
