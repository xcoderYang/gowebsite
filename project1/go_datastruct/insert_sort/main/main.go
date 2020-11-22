package main

import (
	"fmt"
	"sync"
)

func main() {
	// nums := []int{5, 4, 1, 2, 1, 7, 3, 6, 1, 4, 6, 7, 2, 5, 4, 2}
	// nums = insert_sort(nums)
	// fmt.Println(nums)
	sum := 0
	total := 0
	var lock sync.Mutex
	for i := 1; i <= 10; i++ {
		lock.Lock()
		sum += 1
		go func() {
			total += 1
			lock.Unlock()
		}()
	}
	fmt.Println(sum, total)
}

func insert_sort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				temp := nums[j-1]
				nums[j-1] = nums[j]
				nums[j] = temp
			} else {
				break
			}
		}
	}
	return nums
}
