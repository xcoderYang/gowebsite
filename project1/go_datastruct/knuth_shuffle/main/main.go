package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var nums = []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
	}
	nums = knuth_shuffle(nums)
	fmt.Println(nums)
}

func knuth_shuffle(nums []int) []int {
	Len := len(nums)
	for i := Len - 1; i >= 0; i-- {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(i + 1)
		temp := nums[i]
		nums[i] = nums[index]
		nums[index] = temp
	}
	return nums
}
