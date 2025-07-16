package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	grpcInterno "github.com/Ryan-18-system/order-service/internal/order/delivery/grpc" // Import the gRPC server package
	pb "github.com/Ryan-18-system/order-service/internal/order/delivery/grpc/order-service/proto"
	"github.com/Ryan-18-system/order-service/internal/order/delivery/rest"
	"github.com/Ryan-18-system/order-service/internal/order/domain"
	"github.com/Ryan-18-system/order-service/internal/order/repository"
	"github.com/Ryan-18-system/order-service/internal/order/usecase"
	"github.com/go-chi/chi"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// ✅ Cria a tabela se não existir
	err = db.AutoMigrate(&domain.Order{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	var count int64
	db.Model(&domain.Order{}).Count(&count)
	if count == 0 {
		db.Create(&[]domain.Order{
			{CustomerName: "João Silva", Total: 199.90},
			{CustomerName: "Maria Oliveira", Total: 350.50},
			{CustomerName: "Carlos Souza", Total: 75.20},
		})
		log.Println("🚀 Dados de teste inseridos em 'orders'")
	}
	repo := repository.NewPostgresOrderRepository(db)
	uc := usecase.NewOrderUseCase(repo)

	// REST
	r := chi.NewRouter()
	r.Get("/order", rest.GetOrdersHandler(uc))
	go http.ListenAndServe(":8080", r)

	// gRPC
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterOrderServiceServer(s, grpcInterno.NewGrpcServer(uc))
		log.Println("gRPC server listening on :50051")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// GraphQL
	http.HandleFunc("/graphql", rest.GraphQLHandler(uc))
	log.Println("GraphQL server listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
