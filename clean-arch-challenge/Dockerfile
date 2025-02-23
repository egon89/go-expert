# Stage 1: Build the Go application
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /app/ordersystem cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go

# Stage 2: Create a minimal production image
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder --chown=nonroot:nonroot /app/ordersystem /

WORKDIR /

EXPOSE 8000 50051 8080

USER nonroot

CMD ["./ordersystem"]
