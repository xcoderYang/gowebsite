package main

import "fmt"

type node struct {
	Val   interface{}
	Index int
}

func main() {
	nums := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(transpose(nums))
}

func transpose(A [][]int) [][]int {
	nums := make([][]int, len(A))
	for i := range nums {
		nums[i] = make([]int, len(A[i]))
	}
	for i := range A {
		for j := range A[i] {
			nums[j][i] = A[i][j]
		}
	}
	return nums
}
