package tetris2d

/*
Next Fit Decreasing Height (NFDH)

Explanation:
NFDH sorts items by decreasing height and then places them into levels. It starts with the tallest item and continues placing items in the current level until an item doesn't fit. When this happens, it starts a new level.
Pros:

Simple to implement and understand
Fast execution time (O(n log n) due to sorting)
Works well for items with similar heights

Cons:

Can leave significant empty space, especially with varied item sizes
Not optimal for complex arrangements

Use cases:

Quick approximations
Scenarios where speed is more important than optimal packing
When items have similar heights
*/

func NextFitDecreasing(bin Bin, items []Item) {
	// sort items by decreasing height
	levels := []int{}
	currentLevel := 0
	remainingWidth := 0.0

	for _, item := range items {
    if item.Width > remainingWidth {
      // start new level
      currentLevel += 1
      remainingWidth = bin.Width
    } 
	}

	/*
		remaining_width = bin_width

		    for item in items:
		        if item[0] > remaining_width:
		            # Start a new level
		            current_level += 1
		            remaining_width = bin_width

		        if len(levels) <= current_level:
		            levels.append([])

		        levels[current_level].append(item)
		        remaining_width -= item[0]

		    return levels
	*/
}
