package rotateimageinplace

func rotate(matrix [][]int) {
	// transponate matrix by swapping (i,j) and (j,i) elems

	n := len(matrix) // size of square matrix

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// reverse each line
	for i := 0; i < n; i++ {
		first := 0
		last := n - 1
		for first < last {
			matrix[i][first], matrix[i][last] = matrix[i][last], matrix[i][first]
			first += 1
			last -= 1
		}

	}
}
