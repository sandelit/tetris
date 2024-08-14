package tetris2d

func calculateTotalArea(bins []*Bin) float64 {
	var area float64
	for _, bin := range bins {
		area += bin.Width * bin.Height
	}
	return area
}

func calculateUsedArea(bins []*Bin) float64 {
	var area float64
	for _, bin := range bins {
		for _, item := range bin.Items {
			area += item.Width * item.Height
		}
	}
	return area
}

func getResult(bins []*Bin, remainingItems []*Item) BinPackResult {
	result := BinPackResult{
		Bins:           bins,
		RemainingItems: remainingItems,
	}
	result.TotalArea = calculateTotalArea(bins)
	result.UsedArea = calculateUsedArea(bins)
	result.Efficiency = result.UsedArea / result.TotalArea
	return result
}

func (p *Packer) tryToPlaceItem(bin *Bin, item *Item) bool {
	x := 0.0
	y := 0.0
	for _, existingItem := range bin.Items {
		if existingItem.Position.X+existingItem.Width+item.Width <= bin.Width {
			x = existingItem.Position.X + existingItem.Width
			y = existingItem.Position.Y
		} else {
			x = 0
			y = existingItem.Position.Y + existingItem.Height
			if y+item.Height > bin.Height {
				return false
			}
		}
	}
	item.Position = Position{X: x, Y: y}
	bin.Items = append(bin.Items, item)
	return true
}
