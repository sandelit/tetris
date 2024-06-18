package tetris

import (
	"sort"
)

type Box struct {
	Width, Height, Depth int
	X, Y, Z              int
}

type Bin struct {
	Width, Height, Depth, MaxWeight int
}

// Function to sort boxes by volume (width * height * depth) in descending order
func SortBoxesByVolume(boxes []Box) {
	sort.Slice(boxes, func(i, j int) bool {
		vol1 := boxes[i].Width * boxes[i].Height * boxes[i].Depth
		vol2 := boxes[j].Width * boxes[j].Height * boxes[j].Depth
		return vol1 > vol2
	})
}

