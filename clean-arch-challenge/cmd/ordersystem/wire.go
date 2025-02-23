//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/egon89/clean-arch-challenge/internal/entity"
	"github.com/egon89/clean-arch-challenge/internal/event"
	"github.com/egon89/clean-arch-challenge/internal/infra/database"
	"github.com/egon89/clean-arch-challenge/internal/infra/web"
	"github.com/egon89/clean-arch-challenge/internal/usecase"
	"github.com/egon89/clean-arch-challenge/pkg/events"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}

func NewGetAllOrderUseCase(db *sql.DB) *usecase.GetAllOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewGetAllOrderUseCase,
	)
	return &usecase.GetAllOrderUseCase{}
}
