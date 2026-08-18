package calculators

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"gs-1/repositories"
	"testing"
)

// lifted from the config for testing purposes, I'd really mock the interface from the actual service if testing end-to-end
var packs = []repositories.ProductPackSize{
	{Size: 250, Cost: decimal.RequireFromString("10.00")},
	{Size: 500, Cost: decimal.RequireFromString("19.00")},
	{Size: 1000, Cost: decimal.RequireFromString("35.00")},
	{Size: 2000, Cost: decimal.RequireFromString("65.00")},
	{Size: 5000, Cost: decimal.RequireFromString("100.00")},
}

func TestCalculateReturns250ItemWhen1Wanted2(t *testing.T) {

	// Act
	var result = Calculate(1, packs)

	// Assert
	assert.Len(t, result.Packs, 1)
	assert.Equal(t, 250, result.Packs[0].PackSize)
	assert.Equal(t, 1, result.Packs[0].Quantity)
}

func TestCalculateReturns250ItemWhen250Wanted2(t *testing.T) {

	// Act
	var result = Calculate(250, packs)

	// Assert
	assert.Len(t, result.Packs, 1)
	assert.Equal(t, 250, result.Packs[0].PackSize)
	assert.Equal(t, 1, result.Packs[0].Quantity)
}

func TestCalculateReturns500ItemWhen251Wanted2(t *testing.T) {

	// Act
	var result = Calculate(251, packs)

	// Assert
	assert.Len(t, result.Packs, 1)
	assert.Equal(t, 500, result.Packs[0].PackSize)
	assert.Equal(t, 1, result.Packs[0].Quantity)
}

func TestCalculateReturns750ItemWhen501Wanted2(t *testing.T) {

	// Act
	var result = Calculate(501, packs)

	// Assert
	assert.Len(t, result.Packs, 2)
	assert.Equal(t, 500, result.Packs[0].PackSize)
	assert.Equal(t, 1, result.Packs[0].Quantity)
	assert.Equal(t, 250, result.Packs[1].PackSize)
	assert.Equal(t, 1, result.Packs[1].Quantity)
}

func TestCalculateReturns500ItemWhen12001Wanted2(t *testing.T) {

	// Act
	var result = Calculate(12001, packs)

	// Assert
	assert.Len(t, result.Packs, 3)
	assert.Equal(t, 5000, result.Packs[0].PackSize)
	assert.Equal(t, 2, result.Packs[0].Quantity)
	assert.Equal(t, 2000, result.Packs[1].PackSize)
	assert.Equal(t, 1, result.Packs[1].Quantity)
	assert.Equal(t, 250, result.Packs[2].PackSize)
	assert.Equal(t, 1, result.Packs[2].Quantity)
}

func TestCalculateReturns1000ItemWhen751Wanted2(t *testing.T) {

	// Act
	var result = Calculate(751, packs)

	// Assert
	assert.Len(t, result.Packs, 1)
	assert.Equal(t, 1000, result.Packs[0].PackSize)
	assert.Equal(t, 1, result.Packs[0].Quantity)
}

func TestCalculateReturns2PacksWithSmallerLeftOvers2(t *testing.T) {

	// Act
	var result = Calculate(501, packs)

	// Assert
	assert.Len(t, result.Packs, 2)
	assert.Equal(t, 500, result.Packs[0].PackSize)
	assert.Equal(t, 1, result.Packs[0].Quantity)
	assert.Equal(t, 250, result.Packs[1].PackSize)
	assert.Equal(t, 1, result.Packs[1].Quantity)
}

func TestCalculateReturns1PackWhenItFindsMatchingSinglePack2(t *testing.T) {

	// Act
	var result = Calculate(500, []repositories.ProductPackSize{
		{Size: 250, Cost: decimal.RequireFromString("10.00")},
		{Size: 500, Cost: decimal.RequireFromString("19.00")},
	})

	// Assert
	assert.Len(t, result.Packs, 1)
	assert.Equal(t, 500, result.Packs[0].PackSize)
	assert.Equal(t, 1, result.Packs[0].Quantity)
}
