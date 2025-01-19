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

func (uc *GetAllOrderUseCase) Execute() (GetAllOrderOutputDto, error) {
	entities, err := uc.OrderRepository.GetAllOrders()
	if err != nil {
		return GetAllOrderOutputDto{}, err
	}

	var orders []OrderOutputDTO
	for _, entity := range entities {
		order := OrderOutputDTO{
			ID:         entity.ID,
			Price:      entity.Price,
			Tax:        entity.Tax,
			FinalPrice: entity.FinalPrice,
		}
		orders = append(orders, order)
	}

	return GetAllOrderOutputDto{
		Orders: orders,
	}, nil
}
