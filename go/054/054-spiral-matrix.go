func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	top, left := 0, 0
	bottom := len(matrix) - 1 // Total Rows
	right := len(matrix[0]) - 1

	result := make([]int, 0, len(matrix)*len(matrix[0]))

	for top <= bottom && left <= right {
		// right
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// down
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// left
		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		// up
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}
	return result
}
