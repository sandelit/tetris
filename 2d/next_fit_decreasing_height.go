package tetris2d

import (
	"sort"
)

/*
Next Fit Decreasing Height (NFDH)

Explanation:
NFDH sorts items by decreasing height and then places them into levels. It starts with the tallest item and continues placing items in the current level until an item doesn't fit. When this happens, it starts a new level.

Use cases:
Quick approximations
Scenarios where speed is more important than optimal packing
When items have similar heights
*/
func (Packer) NextFitDecreasingHeight(bins []*Bin, items []*Item) (BinPackResult, error) {
	remainingItems := make([]*Item, 0)
	// Sort items by decreasing height
	sort.Slice(items, func(i, j int) bool {
		return items[i].Height > items[j].Height
	})

	for _, item := range items {
		binIndex := -1
		for i, bin := range bins {
			if item.Width <= bin.Width && item.Height <= bin.Height {
				binIndex = i
				break
			}
		}

		if binIndex == -1 {
			// Item does not fit in any bin
			remainingItems = append(remainingItems, item)
			continue
		}

		bin := bins[binIndex]
		x := 0.0
		y := 0.0

		for _, existingItem := range bin.Items {
			if existingItem.Width+item.Width <= bin.Width {
				x = existingItem.Position.X + existingItem.Width
				y = existingItem.Position.Y
			} else {
				x = 0
				y += existingItem.Height
				if y+item.Height > bin.Height {
					binIndex = -1
					break
				}
			}
		}

		if binIndex != -1 {
			item.Position = Position{X: x, Y: y}
			bin.Items = append(bin.Items, item)
		}
	}

	result := BinPackResult{
		Bins:           bins,
		RemainingItems: remainingItems,
	}

	result.TotalArea = calculateTotalArea(bins)
	result.UsedArea = calculateUsedArea(items)
	result.Efficiency = result.UsedArea / result.TotalArea

	return result, nil
}
