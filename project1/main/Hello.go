package main

import (
	"fmt"
	"sort"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var vari = []int{0, 1}
	fmt.Println(largestNumber(vari))
	// fmt.Println("*****************************")
	// fmt.Println(fractionToDecimal(10, 3))
}

type IntSlice []int

func (s IntSlice) Len() int { return len(s) }
func (s IntSlice) Less(i, j int) bool {
	stri := strconv.Itoa(s[i])
	strj := strconv.Itoa(s[j])
	numi := stri + strj
	numj := strj + stri
	return numi > numj
}
func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func largestNumber(nums []int) string {
	sort.Sort(IntSlice(nums))
	S := ""
	index := 0
	for index < len(nums) && nums[index] == 0 {
		index++
	}
	if index == len(nums) {
		return "0"
	}
	for index < len(nums) {
		S = S + strconv.Itoa(nums[index])
		index++
	}

	return S
}
