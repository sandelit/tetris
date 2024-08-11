package tetris2d

type Item struct {
	Name     string
	Width    float64
	Height   float64
	Position Position
}

type Position struct {
	X, Y float64
}

type RotationType int

const (
	RotationType_WH RotationType = iota
	RotationType_HW
)

var RotationTypeStrings = [...]string{
	"RotationType_WH (w,h)",
	"RotationType_HW (h,w)",
}

func (rt RotationType) String() string {
	return RotationTypeStrings[rt]
}

func NewItem(name string, w, h float64) *Item {
	return &Item{
		Name:   name,
		Width:  w,
		Height: h,
	}
}
