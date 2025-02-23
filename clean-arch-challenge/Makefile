run:
	docker-compose up

reset-run:
	docker-compose down -v && docker-compose up

build:
	docker-compose build

dev-run:
	docker-compose -f docker-compose.dev.yaml up

dev-build:
	docker-compose -f docker-compose.dev.yaml build

local-run:
	go run cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go

local-build:
	go build -o ordersystem cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go

clean:
	docker-compose down -v && docker-compose -f docker-compose.dev.yaml down -v
