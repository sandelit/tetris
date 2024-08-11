package tetris2d

import (
	"fmt"
	"strings"
)

type BinPackResult struct {
	Bins           []*Bin
	RemainingItems []*Item
	Efficiency     float64
	TotalArea      float64
	UsedArea       float64
}

func (bpr *BinPackResult) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Packing Efficiency: %.2f%%\n", bpr.Efficiency*100))
	sb.WriteString(fmt.Sprintf("Total Area: %.2f, Used Area: %.2f\n", bpr.TotalArea, bpr.UsedArea))

	var totalPackedItems int
	for _, bin := range bpr.Bins {
		totalPackedItems += len(bin.Items)
	}

	sb.WriteString(fmt.Sprintf("Packed Items: %d, Unpacked Items: %d\n\n", totalPackedItems, len(bpr.RemainingItems)))

	for _, bin := range bpr.Bins {
		sb.WriteString(bin.String())
		sb.WriteString("\n")
	}

	if len(bpr.RemainingItems) > 0 {
		sb.WriteString("\nUnpacked Items:\n")
		for _, item := range bpr.RemainingItems {
			sb.WriteString(fmt.Sprintf("- %s (%.2fx%.2f)\n", item.Name, item.Width, item.Height))
		}
	}

	return sb.String()
}
