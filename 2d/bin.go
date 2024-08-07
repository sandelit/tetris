package tetris2d

import "fmt"

type Bin struct {
	Name   string
	Width  float64
	Height float64

	Items []*Item
}

func NewBin(name string, w, h float64) *Bin {
	return &Bin{
		Name:   name,
		Width:  w,
		Height: h,
		Items:  make([]*Item, 0),
	}
}

func (b *Bin) String() string {
	return fmt.Sprintf("%s(%vx%v)", b.Name, b.Width, b.Height)
}
