package tetris2d

import (
	"fmt"
	"strings"
)

type Bin struct {
	Id     string
	Width  float64
	Height float64

	Items []*Item
}

func NewBin(id string, w, h float64) *Bin {
	return &Bin{
		Id:   id,
		Width:  w,
		Height: h,
		Items:  make([]*Item, 0),
	}
}

func NewBinWithItems(id string, w, h float64, items []*Item) *Bin {
	return &Bin{
		Id:   id,
		Width:  w,
		Height: h,
		Items:  items,
	}
}

func (b *Bin) String() string {
	var itemStrings []string
	for _, item := range b.Items {
		itemStrings = append(itemStrings, fmt.Sprintf("%.2fx%.2f at (%.2f, %.2f)", item.Width, item.Height, item.Position.X, item.Position.Y))
	}

	itemsString := strings.Join(itemStrings, ", ")
	if itemsString == "" {
		itemsString = "empty"
	}

	return fmt.Sprintf("%s (%0.2fx%0.2f): %s", b.Id, b.Width, b.Height, itemsString)
}
