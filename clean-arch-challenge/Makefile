run:
	docker-compose up

reset-run:
	docker-compose down -v && docker-compose up

build:
	docker-compose build

build-force:
	docker-compose build --no-cache

dev-run:
	docker-compose -f docker-compose.dev.yaml up

dev-build:
	docker-compose -f docker-compose.dev.yaml build

dev-build-force:
	docker-compose -f docker-compose.dev.yaml build --no-cache

local-run:
	go run cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go

local-build:
	go build -o ordersystem cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go

clean:
	docker-compose down -v && docker-compose -f docker-compose.dev.yaml down -v
