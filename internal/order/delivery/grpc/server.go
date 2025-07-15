package rest

import (
	"context"

	pb "github.com/Ryan-18-system/order-service/internal/order/delivery/grpc/order-service/proto"

	"github.com/Ryan-18-system/order-service/internal/order/usecase"
)

type grpcServer struct {
	pb.UnimplementedOrderServiceServer
	uc usecase.OrderUseCase
}

func NewGrpcServer(uc usecase.OrderUseCase) pb.OrderServiceServer {
	return &grpcServer{uc: uc}
}

func (s *grpcServer) ListOrders(ctx context.Context, in *pb.Empty) (*pb.OrderListResponse, error) {
	orders, err := s.uc.ListOrders(ctx)
	if err != nil {
		return nil, err
	}

	var resp pb.OrderListResponse
	for _, o := range orders {
		resp.Orders = append(resp.Orders, &pb.Order{
			Id:           int32(o.ID),
			CustomerName: o.CustomerName,
			Total:        o.Total,
			CreatedAt:    o.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return &resp, nil
}
