package usecase

import "github.com/devfullcycle/20-CleanArch/internal/entity"

type OrderOutputDto struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type GetAllOrderOutputDto struct {
	Orders []OrderOutputDTO `json:"orders"`
}

type GetAllOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetAllOrderUseCase(
	orderRepository entity.OrderRepositoryInterface,
) *GetAllOrderUseCase {
	return &GetAllOrderUseCase{
		OrderRepository: orderRepository,
	}
}

func (uc *GetAllOrderUseCase) Execute() GetAllOrderOutputDto {
	order := OrderOutputDTO{
		ID:         "123",
		Price:      10,
		Tax:        1.99,
		FinalPrice: 11.99,
	}

	orders := []OrderOutputDTO{order}

	return GetAllOrderOutputDto{
		Orders: orders,
	}
}
