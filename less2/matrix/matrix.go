package matrix

import (
	"math/rand"
	"time"
)

func GenerateMatrix(rows, cols int) [][]int {

	rand.Seed(time.Now().UnixNano())

	size := rows * cols
	numbers := rand.Perm(size)

	matrix := make([][]int, rows)

	k := 0

	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = numbers[k]
			k++
		}
	}

	return matrix
}
