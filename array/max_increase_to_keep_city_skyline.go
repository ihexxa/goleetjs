package array

func maxIncreaseKeepingSkyline(grid [][]int) int {
	xHighest := make([]int, len(grid))
	yHighest := make([]int, len(grid))

	for i, row := range grid {
		for j, cell := range row {
			if cell > xHighest[i] {
				xHighest[i] = cell
			}

			if cell > yHighest[j] {
				yHighest[j] = cell
			}
		}
	}

	sum := 0
	for i, row := range grid {
		for j, cell := range row {
			if xHighest[i] < yHighest[j] {
				sum += xHighest[i] - cell
			} else {
				sum += yHighest[j] - cell
			}
		}
	}

	return sum
}
