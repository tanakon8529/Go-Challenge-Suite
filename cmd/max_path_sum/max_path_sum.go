package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxPathSum(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	// Start from the second to last row and move upwards
	for row := len(triangle) - 2; row >= 0; row-- {
		for col := 0; col < len(triangle[row]); col++ {
			// Modify the current row by adding the maximum of the values below
			triangle[row][col] += max(triangle[row+1][col], triangle[row+1][col+1])
		}
	}
	return triangle[0][0]
}
