package application

import (
	"math/rand"
	"time"
)

type Percolation struct {
	matrix [][]int
}

func Percolation_init(n int, p float64) Percolation {
	var per Percolation
	rand.Seed(time.Now().UnixNano())
	per.matrix = make([][]int, n)
	for i, _ := range per.matrix {
		per.matrix[i] = make([]int, n)
		for j, _ := range per.matrix[i] {
			if rand.Float64() < p {
				per.matrix[i][j] = 0
			} else {
				per.matrix[i][j] = 1
			}
		}
	}
	return per
}
func (per Percolation) GetMatrix() [][]int {
	return per.matrix
}
