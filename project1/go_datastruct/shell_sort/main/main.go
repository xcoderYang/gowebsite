package main

import "fmt"

func main() {
	nums := []int{12, 5, 5, 1, 13, 8, 12, 5, 16, 19, 15, 12, 20, 19, 23, 1, 5}
	nums = shell_sort(nums)
	fmt.Println(nums)
}

func shell_sort(nums []int) []int {
	N := len(nums)
	h := 1
	for h < N/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j >= h && nums[j] < nums[j-h]; j -= h {
				temp := nums[j]
				nums[j] = nums[j-h]
				nums[j-h] = temp
			}
		}
		h /= 3
	}
	return nums
}
