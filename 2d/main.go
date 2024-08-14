package tetris2d

import "fmt"

func main() {
	packer := NewPacker()

	bins := []*Bin{
		{Width: 10, Height: 10},
		{Width: 5, Height: 5},
	}

	items := []*Item{
		{Width: 2, Height: 2},
		{Width: 3, Height: 3},
		{Width: 5, Height: 5},
	}

	result := packer.NextFitDecreasingHeight(bins, items)

	fmt.Println("Packed bins:", result)
}
