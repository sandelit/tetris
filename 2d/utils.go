package tetris2d

func calculateTotalArea(bins []*Bin) float64 {
	var area float64
	for _, bin := range bins {
		area += bin.Width * bin.Height
	}
	return area
}

func calculateUsedArea(items []*Item) float64 {
	var area float64
	for _, item := range items {
		area += item.Width * item.Height
	}
	return area
}
