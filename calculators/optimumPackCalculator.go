package calculators

import (
	"gs-1/repositories"
	"sort"
)

type PackCalculationResult struct {
	Packs []repositories.PackAllocation
}

func Calculate(desiredQuantity int, packSizes []repositories.ProductPackSize) PackCalculationResult {

	// The largest we sell
	var largestPackSize = packSizes[0].Size
	for _, pack := range packSizes[1:] {
		if pack.Size > largestPackSize {
			largestPackSize = pack.Size
		}
	}

	// Getting one largest
	var maximumTotalToCheck = desiredQuantity + largestPackSize - 1

	var maxAllowed = 99999 // From validation elsewhere in application (would be config driven in a proper application)
	var packsForTotal = make([]int, maximumTotalToCheck+1)
	var previousPack = make([]int, maximumTotalToCheck+1)

	for i := range packsForTotal {
		packsForTotal[i] = maxAllowed
	}
	packsForTotal[0] = 0

	for total := 1; total <= maximumTotalToCheck; total++ {
		for _, pack := range packSizes {
			if total >= pack.Size &&
				packsForTotal[total-pack.Size] != maxAllowed {

				var currCount = packsForTotal[total-pack.Size] + 1

				if currCount < packsForTotal[total] {
					packsForTotal[total] = currCount
					previousPack[total] = pack.Size
				}
			}
		}
	}

	purchasedQuantity := 0
	for total := desiredQuantity; total <= maximumTotalToCheck; total++ {
		if packsForTotal[total] != maxAllowed {
			purchasedQuantity = total
			break
		}
	}

	// Add up our totals (quantity) for each pack
	var selectedPacks = make(map[int]int) // basically a c# dict
	for total := purchasedQuantity; total > 0; {
		packSize := previousPack[total]
		selectedPacks[packSize]++
		total -= packSize
	}

	var allocations []repositories.PackAllocation
	//var allocations = make([]repositories.PackAllocation, 0, len(selectedPacks))
	for size := range selectedPacks {
		// Go back to our original packs and fill out the allocations
		var found = FindInPacks(size, packSizes)
		var allocation = repositories.PackAllocation{
			Quantity: selectedPacks[size],
			PackCost: found.Cost,
			PackSize: found.Size,
		}
		allocations = append(allocations, allocation)
	}

	// Sort the allocations (size desc order)
	sort.Slice(allocations, func(i, j int) bool {
		return allocations[i].PackSize > allocations[j].PackSize
	})

	return PackCalculationResult{Packs: allocations}
}

func FindInPacks(size int, packSizes []repositories.ProductPackSize) repositories.ProductPackSize {
	for _, pack := range packSizes {
		if pack.Size == size {
			return pack
		}
	}

	// Should never get here (probably something better to do/return)
	return repositories.ProductPackSize{}
}
