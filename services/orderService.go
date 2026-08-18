package services

import (
	"errors"
	"gs-1/calculators"
	"gs-1/repositories"
	"math/rand/v2"

	"github.com/shopspring/decimal"
)

var ErrQuantityMustBePositive = errors.New("quantity must be positive")

type NewOrder struct { // Not Joy Division

	ProductId int
	Quantity  int
}

type OrderService interface {
	Create(newOrder NewOrder) (*repositories.OrderSummary, error)
}

type orderService struct {
	repo        *repositories.OrderRepository
	productRepo *repositories.ProductRepository
}

func NewOrderService(repo *repositories.OrderRepository,
	productRepo *repositories.ProductRepository) OrderService {
	return &orderService{
		repo:        repo,
		productRepo: productRepo,
	}
}

func (o *orderService) Create(newOrder NewOrder) (*repositories.OrderSummary, error) {

	if newOrder.Quantity < 1 {
		return nil, ErrQuantityMustBePositive
	}

	product, err := o.productRepo.Get(newOrder.ProductId)
	if err != nil {
		return nil, err
	}
	var packs []repositories.PackAllocation

	if len(product.PackSizes) > 0 {
		// Work out cheapest way to buy desired quantity
		packs = calculators.Calculate(newOrder.Quantity, product.PackSizes).Packs
	}

	// Sum up the actual quantity the user will receive to send back to UI
	actualQuantity := 0
	for _, pack := range packs {
		actualQuantity += pack.PackSize * pack.Quantity
	}

	return &repositories.OrderSummary{
		OrderId:         rand.IntN(9999),
		ProductId:       newOrder.ProductId,
		PackAllocation:  packs,
		DesiredQuantity: newOrder.Quantity,
		ActualQuantity:  actualQuantity,
		Cost:            getTotalCost(packs, product),
	}, nil
}

// Retrieves the total cost of all packs and quantities
func getTotalCost(packs []repositories.PackAllocation, product *repositories.Product) decimal.Decimal {
	var totalCost = decimal.NewFromInt(0)
	for _, allocation := range packs {
		for _, pack := range product.PackSizes {
			if pack.Size == allocation.PackSize {
				totalCost = totalCost.Add(pack.Cost.Mul(decimal.NewFromInt(int64(allocation.Quantity)))) // Not sure if this is idiomatic Go
				break
			}
		}
	}
	return totalCost
}
