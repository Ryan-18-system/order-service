package rest

import (
	"encoding/json"
	"net/http"

	"github.com/Ryan-18-system/order-service/internal/order/usecase"
)

func GetOrdersHandler(uc usecase.OrderUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orders, err := uc.ListOrders(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(orders)
	}
}
