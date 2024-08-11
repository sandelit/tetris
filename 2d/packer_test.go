package tetris2d

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestPacker(t *testing.T) {
	packer := NewPacker()

	bin1 := NewBin(uuid.NewString(), 10.0, 10.0)
	item1 := NewItem(uuid.NewString(), 10.0, 10.0)
	bin1WithItems := bin1
	bin1WithItems.Items = append(bin1WithItems.Items, item1)

	tests := []struct {
		name     string
		bins     []*Bin
		items    []*Item
		expected BinPackResult
	}{
		{
			name:  "Item fits in bin",
			bins:  []*Bin{bin1},
			items: []*Item{item1},
			expected: BinPackResult{
				Bins:           []*Bin{bin1WithItems},
				RemainingItems: []*Item{},
				Efficiency:     1.0,
				TotalArea:      100.0,
				UsedArea:       100.0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := packer.NextFitDecreasingHeight(tt.bins, tt.items)
			t.Logf("Got result: %#v", got)
			t.Logf("Want result: %#v", tt.expected)
			compareBinPackResults(t, got, tt.expected)
		})
	}
}

func compareBinPackResults(t *testing.T, got, want BinPackResult) {
	if len(got.Bins) != len(want.Bins) {
		t.Errorf("Bin count mismatch: got %d, want %d", len(got.Bins), len(want.Bins))
		return
	}

	for i, gotBin := range got.Bins {
		wantBin := want.Bins[i]
		if gotBin.Id != wantBin.Id {
			t.Errorf("Bin[%d] name mismatch: got %s, want %s", i, gotBin.Id, wantBin.Id)
		}
		if gotBin.Width != wantBin.Width {
			t.Errorf("Bin[%d] width mismatch: got %f, want %f", i, gotBin.Width, wantBin.Width)
		}
		if gotBin.Height != wantBin.Height {
			t.Errorf("Bin[%d] height mismatch: got %f, want %f", i, gotBin.Height, wantBin.Height)
		}
		// Compare items in the bin...
	}

	// Compare other fields...
	if !reflect.DeepEqual(got.RemainingItems, want.RemainingItems) {
		t.Errorf("RemainingItems mismatch: got %#v, want %#v", got.RemainingItems, want.RemainingItems)
	}

	if got.Efficiency != want.Efficiency {
		t.Errorf("Efficiency mismatch: got %f, want %f", got.Efficiency, want.Efficiency)
	}

	if got.TotalArea != want.TotalArea {
		t.Errorf("TotalArea mismatch: got %f, want %f", got.TotalArea, want.TotalArea)
	}

	if got.UsedArea != want.UsedArea {
		t.Errorf("UsedArea mismatch: got %f, want %f", got.UsedArea, want.UsedArea)
	}
}
