package main

import "fmt"

func main() {
	nums := "MERGESORTEXAMPLE"
	fmt.Println(merge_sort(0, len(nums)-1, nums))
}

func merge_sort(start int, end int, nums string) string {
	if start == end {
		return string(nums[start])
	}
	mid := (start + end) / 2
	str1 := merge_sort(start, mid, nums)
	str2 := merge_sort(mid+1, end, nums)
	return merge(str1, str2)
}

func merge(str1 string, str2 string) string {
	newString := ""
	i := 0
	j := 0
	for i < len(str1) && j < len(str2) {
		if str1[i] < str2[j] {
			newString += string(str1[i])
			i++
		} else {
			newString += string(str2[j])
			j++
		}
	}
	if i == len(str1) {
		newString += string(str2[j:])
	}
	if j == len(str2) {
		newString += string(str1[i:])
	}
	return newString
}
