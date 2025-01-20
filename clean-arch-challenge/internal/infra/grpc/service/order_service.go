package service

import (
	"context"

	"github.com/devfullcycle/20-CleanArch/internal/infra/grpc/pb"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	GetAllOrderUseCase usecase.GetAllOrderUseCase
}

func NewOrderService(
	createOrderUseCase usecase.CreateOrderUseCase,
	getAllOrderUseCase usecase.GetAllOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		GetAllOrderUseCase: getAllOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) GetAllOrders(ctx context.Context, in *pb.Blank) (*pb.OrderList, error) {
	output, err := s.GetAllOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var orders []*pb.Order
	for _, orderDto := range output.Orders {
		order := &pb.Order{
			Id:         orderDto.ID,
			Price:      float32(orderDto.Price),
			Tax:        float32(orderDto.Tax),
			FinalPrice: float32(orderDto.FinalPrice),
		}
		orders = append(orders, order)
	}

	return &pb.OrderList{
		Orders: orders,
	}, nil
}
