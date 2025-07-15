package domain

import "context"

type OrderRepository interface {
    List(ctx context.Context) ([]Order, error)
}