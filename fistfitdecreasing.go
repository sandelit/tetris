package tetris

import "fmt"

// Function to perform First-Fit Decreasing packing of boxes into a bin
func FirstFitDecreasing(boxes []Box, bin Bin) ([]Box, error) {
	SortBoxesByVolume(boxes)

	var filledSpace []Box

	// Initialize space in the bin
	remainingSpace := []Box{
		{bin.Width, bin.Height, bin.Depth, 0, 0, 0},
	}

	for _, box := range boxes {
		placed := false
		for i := 0; i < len(remainingSpace); i++ {
			space := remainingSpace[i]
			if box.Width <= space.Width && box.Height <= space.Height && box.Depth <= space.Depth {
				box.X, box.Y, box.Z = space.X, space.Y, space.Z
				filledSpace = append(filledSpace, box)
				remainingSpace = append(remainingSpace[:i], remainingSpace[i+1:]...)
				
				// Update remaining space
				if box.Width < space.Width {
					remainingSpace = append(remainingSpace, Box{space.Width - box.Width, space.Height, space.Depth, space.X + box.Width, space.Y, space.Z})
				}
				if box.Height < space.Height {
					remainingSpace = append(remainingSpace, Box{box.Width, space.Height - box.Height, space.Depth, space.X, space.Y + box.Height, space.Z})
				}
				if box.Depth < space.Depth {
					remainingSpace = append(remainingSpace, Box{box.Width, box.Height, space.Depth - box.Depth, space.X, space.Y, space.Z + box.Depth})
				}
				
				placed = true
				break
			}
		}
		if !placed {
			return nil, fmt.Errorf("not all boxes fit in the bin")
		}
	}

	return filledSpace, nil
}
