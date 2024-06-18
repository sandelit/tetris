package tetris

import (
	"testing"
)

func TestFirstFitDecreasing(t *testing.T) {
	tests := []struct {
		bin     Bin
		boxes   []Box
		success bool
	}{
		{
			bin: Bin{Width: 5, Height: 4, Depth: 4},
			boxes: []Box{
				{Width: 2, Height: 2, Depth: 2},
				{Width: 1, Height: 1, Depth: 1},
				{Width: 1, Height: 1, Depth: 1},
				{Width: 2, Height: 1, Depth: 1},
				{Width: 2, Height: 1, Depth: 1},
			},
			success: true,
		},
		{
			bin: Bin{Width: 5, Height: 4, Depth: 4},
			boxes: []Box{
				{Width: 3, Height: 3, Depth: 3},
				{Width: 3, Height: 3, Depth: 3},
			},
			success: false,
		},
		{
			bin: Bin{Width: 10, Height: 10, Depth: 10},
			boxes: []Box{
				{Width: 5, Height: 5, Depth: 5},
				{Width: 5, Height: 5, Depth: 5},
				{Width: 5, Height: 5, Depth: 5},
				{Width: 5, Height: 5, Depth: 5},
			},
			success: true,
		},
		{
			bin: Bin{Width: 2, Height: 2, Depth: 2},
			boxes: []Box{
				{Width: 3, Height: 3, Depth: 3},
			},
			success: false,
		},
	}

	for _, tt := range tests {
		packedBoxes, err := FirstFitDecreasing(tt.boxes, tt.bin)
		if tt.success && err != nil {
			t.Errorf("expected success but got error: %v", err)
		}
		if !tt.success && err == nil {
			t.Errorf("expected error but got success: %v", packedBoxes)
		}
	}
}
