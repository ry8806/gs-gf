package repositories

import "github.com/shopspring/decimal"

type PackAllocation struct {
	PackSize int             `json:"packSize"` // How many is in each pack
	Quantity int             `json:"quantity"` // The Number of packs
	PackCost decimal.Decimal `json:"packCost"` // The cost of each individual pack
}

type OrderSummary struct {
	OrderId         int              `json:"orderId"`
	ProductId       int              `json:"productId"`
	DesiredQuantity int              `json:"desiredQuantity"`
	ActualQuantity  int              `json:"actualQuantity"`
	Cost            decimal.Decimal  `json:"cost"`
	PackAllocation  []PackAllocation `json:"packAllocation"`
}

type OrderCreator interface {
	Create(order OrderSummary) (bool, error)
}

type OrderRepository struct {
}

func (o *OrderRepository) Create(order OrderSummary) (bool, error) {
	// TODO: Save to a Db
	return true, nil
}
