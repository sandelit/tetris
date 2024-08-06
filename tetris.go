package tetris

type Packer struct{}

func NewPacker() *Packer {
	return &Packer{}
}

type Bin struct {
	id        string
	width     float64
	height    float64
	depth     float64
	maxWeight float64
	volume    float64

	// items currently in the bin
	items []*Item

	// longest side of the bin
	maxLength float64

	// total volume of items in the bin
	itemsVolume float64

	// total weight of items in the bin
	itemsWeight float64
}

type Item struct {
	id     string
	width  float64
	height float64
	depth  float64
	weight float64
	volume float64

	maxLength float64
	rotation  Rotation
	position  Position
}

type Rotation int

const (
	WHD = iota
	WDH
	HWD
	HDW
	DHW
	DWH
)

var rotationName = map[Rotation]string{
	WHD: "WHD",
	WDH: "WDH",
	HWD: "HWD",
	HDW: "HDW",
	DHW: "DHW",
	DWH: "DWH",
}

func (rotation Rotation) String() string {
	return rotationName[rotation]
}

type Position struct {
	x, y, z float64
}
