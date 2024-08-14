package tetris2d

import (
	"sort"
)

/*
Next Fit Decreasing Height (NFDH)

NFDH sorts items by decreasing height and then places them into levels. It starts with the tallest item and continues placing items in the current level until an item doesn't fit. When this happens, it starts a new level.

Use cases:
- Quick approximations
- Scenarios where speed is more important than optimal packing
- When items have similar heights
*/

func (p *Packer) NextFitDecreasingHeight(bins []*Bin, items []*Item) BinPackResult {
	// Sort items by decreasing height
	sort.Slice(items, func(i, j int) bool {
		return items[i].Height > items[j].Height
	})

	remainingItems := make([]*Item, 0)
	for _, item := range items {
		placed := false
		for _, bin := range bins {
			if item.Width <= bin.Width && item.Height <= bin.Height {
				if p.tryPlaceItem(bin, item) {
					placed = true
					break
				}
			}
		}
		if !placed {
			remainingItems = append(remainingItems, item)
		}
	}

	return getResult(bins, remainingItems)
}

