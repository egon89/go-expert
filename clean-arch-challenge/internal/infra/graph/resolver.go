package graph

import "github.com/egon89/clean-arch-challenge/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
	GetAllOrderUseCase usecase.GetAllOrderUseCase
}
