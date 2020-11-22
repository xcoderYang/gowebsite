package main

import (
	"fmt"
	"sort"
)

type MyCalendar struct {
	cache []map[string]int
}

func Constructor() MyCalendar {
	var calender MyCalendar
	return calender
}

func (this *MyCalendar) Book(start int, end int) bool {
	todo := make(map[string]int)
	todo["start"] = start
	todo["end"] = end
	if len(this.cache) <= 0 {
		this.cache = append(this.cache, todo)
		return true
	}
	todolist := make([]map[string]int, 0)
	todolist = append(todolist, todo)
	if end <= this.cache[0]["start"] {
		this.cache = append(todolist, this.cache...)
		return true
	} else if start >= this.cache[len(this.cache)-1]["end"] {
		this.cache = append(this.cache, todo)
		return true
	}
	for i := 1; i < len(this.cache); i++ {
		if todo["end"] <= this.cache[i]["start"] && todo["start"] >= this.cache[i-1]["end"] {
			fmt.Println(i)
			this.cache = append(this.cache[:i], append([]map[string]int{todo}, this.cache[i:]...)...)
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}

type topNode struct {
	Val  int
	Next *topNode
}

func findRepeatedDnaSequences(s string) []string {
	strMap := make(map[string]int)
	ans := make([]string, 0)
	for i := 0; i <= len(s)-10; i++ {
		j := i + 9
		str := s[i : j+1]
		if val, ok := strMap[str]; ok {
			if val > 1 {
				break
			}
			tmpStr := make([]byte, 10)
			copy(tmpStr, str)
			ans = append(ans, string(tmpStr))
			strMap[str]++
		} else {
			strMap[str] = 1
		}
	}
	return ans
}

// func findSubstringInWraproundString(p string) int {
// 	s := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
// 	ans:=0
// 	for i:=0; i<len(p); i++{

// 	}
// }

// func new21Game(N int, K int, W int) float64 {
// 	var P float64
// 	for x := K; x <= N; x++ {
// 		for i:=K;i<=x;i++{

// 		}
// 		var px float64 = 0.0
// 		px += float64(x-K) / float64(W)
// 		remain:=
// 	}
// }

// func lengthOfLIS(nums []int) int {
// 	cache := []int{nums[0]}
// 	for i := 0; i < len(nums); i++ {

// 	}
// }

// func findP(cache []int, goal int) int {

// }

func minFlipsMonoIncr(S string) int {
	preSum := []int{}
	if S[0] == '1' {
		preSum = append(preSum, 1)
	} else {
		preSum = append(preSum, 0)
	}
	// i以及i以后的数全为1
	for i := 1; i < len(S); i++ {
		if S[i] == '1' {
			preSum = append(preSum, preSum[i-1]+1)
		} else {
			preSum = append(preSum, preSum[i-1])
		}
	}
	totalSum := preSum[len(preSum)-1]
	minCount := len(S) - totalSum
	for i := 0; i < len(S); i++ {
		count := preSum[i] + (len(S) - 1 - i - totalSum + preSum[i])
		if minCount > count {
			minCount = count
		}
	}
	return minCount
}

func findMinFibonacciNumbers(k int) int {
	ans := 0
	fibo := []int{1, 1}
	for i := 2; i < 50; i++ {
		fibo = append(fibo, fibo[i-2]+fibo[i-1])
	}
	for i := 49; i >= 0; i-- {
		if fibo[i] <= k {
			k -= fibo[i]
			ans++
		}
		if k == 0 {
			break
		}
	}
	return ans
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums, 0, len(nums)-k)
	fmt.Println(nums)
	reverse(nums, len(nums)-k, len(nums))
	fmt.Println(nums)
	reverse(nums, 0, len(nums))
	fmt.Println(nums)
}

// func reverse(arr []int, start int, end int){
// 	for i:=start; i<start+(end-start)/2; i++{

// 	}
// }

func getLastMoment(n int, left []int, right []int) int {
	var (
		leftMax  int
		rightMin int
	)
	if len(left) > 0 {
		leftMax = getMax(left)
	} else {
		leftMax = 0
	}
	if len(right) > 0 {
		rightMin = getMin(right)
	} else {
		rightMin = n
	}
	if leftMax > (n - rightMin) {
		return leftMax
	} else {
		return n - rightMin
	}
}

func getMin(arr []int) int {
	min := arr[0]
	for i := 0; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}
	return min
}
func getMax(arr []int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}
func f1() int {
	r := 10
	defer func() {
		r++
	}()
	return r
}

func maxAliveYear(birth []int, death []int) int {
	yearCount := [101]int{0}
	for i := 0; i < len(birth); i++ {
		for j := birth[i]; j <= death[i]; j++ {
			yearCount[j-1900]++
		}
	}
	max := 0
	maxIndex := 0
	for i := 0; i < len(yearCount); i++ {
		if max < yearCount[i] {
			maxIndex = i
			max = yearCount[i]
		}
	}
	return maxIndex + 1900
}

func numTeams(rating []int) int {
	var (
		LEN = len(rating)
		i   = 0
		ans = 0
	)
	for j := 1; j < LEN-1; j++ {
		large := 0
		small := 0
		for i = j; i < LEN; i++ {
			if rating[i] > rating[j] {
				large++
			} else if rating[i] < rating[j] {
				small++
			}
		}
		for i = 0; i < j; i++ {
			if rating[i] < rating[j] {
				ans += large
			} else if rating[i] > rating[j] {
				ans += small
			}
		}
	}
	return ans
}

func findUnsortedSubarray(nums []int) int {
	var (
		i        int
		j        int
		localMax int
		localMin int
		LEN      = len(nums)
	)
	for i = 0; i < LEN-1; i++ {
		if nums[i] > nums[i+1] {
			break
		}
	}
	localMin = nums[i]

	for ; i < LEN; i++ {
		if localMin > nums[i] {
			localMin = nums[i]
		}
	}
	for j = LEN - 1; j > 0; j-- {
		if nums[j] < nums[j-1] {
			break
		}
	}
	localMax = nums[j]
	for ; j >= 0; j-- {
		if localMax < nums[j] {
			localMax = nums[j]
		}
	}
	start := 0
	end := 0
	for i = 0; i < LEN; i++ {
		if nums[i] > localMin {
			start = i
			break
		}
	}
	for j = LEN - 1; j >= 0; j-- {
		if nums[j] < localMax {
			end = j
			break
		}
	}
	if start == end {
		return 0
	}
	return end - start + 1
}

func findDuplicates(nums []int) []int {
	ans := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		var index int
		if nums[i] < 0 {
			index = -nums[i] - 1
		} else {
			index = nums[i] - 1
		}
		if nums[index] < 0 {
			ans = append(ans, index+1)
		} else {
			nums[index] *= -1
		}
	}
	return ans
}

func prevPermOpt1(A []int) []int {
	for i := len(A) - 1; i >= 0; i-- {
		cache := i + 1
		for j := i + 1; j < len(A); j++ {
			if A[j] < A[i] && A[j] > A[cache] {
				cache = j
			}
		}
		if cache < len(A) && A[cache] < A[i] {
			A[cache], A[i] = A[i], A[cache]
			break
		}
	}
	return A
}

func smallestDifference(a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	minNum := different(a[0], b[0])
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		diff := different(a[i], b[j])
		if minNum > diff {
			minNum = diff
		}
		if a[i] > b[j] {
			j++
		} else {
			i++
		}
	}
	return minNum
}

func different(a int, b int) int {
	ans := a - b
	if a < b {
		return -ans
	} else {
		return ans
	}
}

func findDiagonalOrder(nums [][]int) []int {
	cache := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			if len(cache) < i+j+1 {
				arr := make([]int, 0)
				cache = append(cache, arr)
			}
			cache[i+j] = append(cache[i+j], nums[i][j])
		}
	}
	ans := make([]int, 0)
	for i := 0; i < len(cache); i++ {
		for j := len(cache[i]) - 1; j >= 0; j-- {
			ans = append(ans, cache[i][j])
		}
	}
	return ans
}

func test(arr []int) {
	start := 0
	end := 9
	for i := start; i < end/2; i++ {
		temp := arr[i]
		arr[i] = arr[end-1-i+start]
		arr[end-1-i+start] = temp
	}
}

func pancakeSort(arr []int) []int {
	LEN := len(arr)
	ans := make([]int, 0, 1000)
	for i := LEN; i > 0; i-- {
		ans = iteration(arr[:i], ans, i, true)
	}
	return ans
}

func iteration(arr []int, ans []int, num int, direction bool) []int {
	LEN := len(arr)
	// 数字在指定位置
	if direction && arr[LEN-1] == num || !direction && arr[0] == num {
		return ans
		// 数字不在指定位置，但是通过一次转换即可
	} else if direction && arr[0] == num || !direction && arr[LEN-1] == num {
		ans = append(ans, len(arr))
		reverse(arr[:], 0, LEN)
		// 数组不在指定位置，且一次转换无法成功
	} else if direction {
		ans = iteration(arr[:LEN-1], ans, num, false)
		reverse(arr[:], 0, LEN)
		ans = append(ans, LEN)
	} else {
		ans = iteration(arr[:LEN-1], ans, num, false)
	}
	return ans
}

func reverse(arr []int, start int, end int) {
	for i := start; i < start+(end-start)/2; i++ {
		temp := arr[i]
		arr[i] = arr[end-1-i+start]
		arr[end-1-i+start] = temp
	}
}
