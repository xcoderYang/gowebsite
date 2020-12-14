package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

type MyCalendar struct {
	cache []map[string]int
}

//func Constructor() MyCalendar {
//	var calender MyCalendar
//	return calender
//}

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
	//gameOfLife([][]int{
	//	{0,1,0},
	//	{0,0,1},
	//	{1,1,1},
	//	{0,0,0},
	//})
	//fmt.Println(getHint("1123","0111"))
	//fmt.Println(maximumSwap(1993))
	//fmt.Println(math.pi)
	fmt.Println(canFinish(4,[][]int{
		{3,2},
		{2,1},
		{1,0},
		{0,3},
	}))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	course:=make([][]int, numCourses)
	for i:=0; i<numCourses; i++{
		course[i] = make([]int,0)
	}
	for i := 0; i<len(prerequisites); i++{
		// to   表示目的
		// from 表示先决条件
		to := prerequisites[i][0]
		from := prerequisites[i][1]
		course[from] = append(course[from], to)
	}
	for{
		node:=-1
		for i:=0; i<numCourses; i++{
			if len(course[i]) == 0{
				node = i
				course[i] = append(course[i], -1)
				break
			}else if course[i][0] == -1{
				continue
			}
		}
		if node == -1{	
			break
		}
		for i:=0; i<numCourses; i++{
			j:=0
			for j<len(course[i]){
				if course[i][j] == node{
					course[i] = append(course[i][:j], course[i][j+1:]...)
					break
				}
				j++
			}
		}
	}
	for i:=0; i<numCourses; i++{
		if len(course[i])>1 || course[i][0] != -1{
			return false
		}
	}
	return true
}

func uniquePaths(m int, n int) int {
	dp:=make([][]int, m)
	for i:=0; i<m; i++ {
		dp[i] = make([]int, n)
		dp[0][0] = 1
		for j:=0; j<n; j++{
			if i == 0 && j == 0{
				continue
			}
			if i == 0{
				dp[i][j] = dp[i][j-1]
			}else if j == 0{
				dp[i][j] = dp[i-1][j]
			}else{
				dp[i][j] = dp[i-1][j]+dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func maximumSwap(num int) int {
	temp := []byte(strconv.Itoa(num))
	maxStack := make([]map[string]int, len(temp))

	maxi := len(temp) - 1
	max := temp[maxi]
	maxStack[maxi] = make(map[string]int)
	maxStack[maxi]["val"] = int(max) - int('0')
	maxStack[maxi]["key"] = maxi

	for i := len(temp) - 2; i >= 0; i-- {
		maxStack[i] = make(map[string]int)
		if temp[i] > max {
			maxStack[i]["val"] = int(temp[i] - '0')
			maxStack[i]["key"] = i
			max = temp[i]
		} else {
			maxStack[i]["val"] = maxStack[i+1]["val"]
			maxStack[i]["key"] = maxStack[i+1]["key"]
		}
	}
	fmt.Println(maxStack)
	for i := 0; i < len(temp); i++ {
		if int(temp[i]-'0') < maxStack[i]["val"] {
			tp := temp[i]
			temp[i] = uint8(maxStack[i]["val"]) + '0'
			temp[maxStack[i]["key"]] = tp
			break
		}
	}
	ans, _ := strconv.Atoi(string(temp))
	return ans
	//fmt.Println(string(temp))

}

type NumArray struct {
	Cache       []int
	UpdateCache []int
}

func Constructor(nums []int) NumArray {
	NA := NumArray{}
	if len(nums) == 0 {
		return NA
	}
	NA.UpdateCache = make([]int, len(nums))
	NA.Cache = make([]int, len(nums))
	NA.Cache[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		NA.Cache[i] = NA.Cache[i-1] + nums[i]
	}
	return NA
}

func (this *NumArray) Update(i int, val int) {
	if len(this.Cache) == 0 {
		return
	}
	var diff int
	if i != 0 {
		diff = this.Cache[i] - this.Cache[i-1] - val
	} else {
		diff = this.Cache[i] - val
	}
	this.UpdateCache[i] = val
	for k := i; k < len(this.Cache); k++ {
		this.Cache[k] = this.Cache[k] - diff
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	if len(this.Cache) == 0 {
		return 0
	}
	if i == 0 {
		return this.Cache[j]
	} else {
		return this.Cache[j] - this.Cache[i]
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */

type NumMatrix struct {
	Cache [][]int
}

//func Constructor(matrix [][]int) NumMatrix {
//	NM:=NumMatrix{}
//	NM.Cache = make([][]int, len(matrix))
//	for i:=0; i<len(matrix); i++{
//		NM.Cache[i] = make([]int, len(matrix[0]))
//		NM.Cache[0][0] = matrix[0][0]
//		for j:=0; j<len(matrix[0]); j++{
//			if i==0 && j==0{
//				continue
//			}
//			if i==0{
//				NM.Cache[i][j] = NM.Cache[i][j-1]+matrix[i][j]
//			}else if j==0{
//				NM.Cache[i][j] = NM.Cache[i-1][j]+matrix[i][j]
//			}else{
//				NM.Cache[i][j] = NM.Cache[i-1][j]+NM.Cache[i][j-1]-NM.Cache[i-1][j-1]+matrix[i][j]
//			}
//		}
//	}
//	return NM
//}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 == 0 && col1 == 0 {
		return this.Cache[row2][col2]
	}
	if row1 == 0 {
		return this.Cache[row2][col2] - this.Cache[row2][col1-1]
	}
	if col1 == 0 {
		return this.Cache[row2][col2] - this.Cache[row1-1][col2]
	}
	return this.Cache[row2][col2] - this.Cache[row2][col1-1] - this.Cache[row1-1][col2] + this.Cache[row1-1][col1-1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

func matrixScore(A [][]int) int {
	for i := 0; i < len(A); i++ {
		if A[i][0] == 0 {
			for j := 0; j < len(A[i]); j++ {
				A[i][j] = 1 - A[i][j]
			}
		}
	}

	for j := 0; j < len(A[0]); j++ {
		count1, count0 := 0, 0
		for i := 0; i < len(A); i++ {
			if A[i][j] == 1 {
				count1++
			} else {
				count0++
			}
		}
		if count0 > count1 {
			for i := 0; i < len(A); i++ {
				A[i][j] = 1 - A[i][j]
			}
		}
	}
	count := 0
	for i := 0; i < len(A); i++ {
		carry := 0.0
		for j := len(A[i]) - 1; j >= 0; j-- {
			count += A[i][j] * int(math.Pow(2.0, carry))
			carry++
		}
	}
	return count
}

func getHint(secret string, guess string) string {
	secretChar := []byte(secret)
	guessChar := []byte(guess)
	Bull, Cow := 0, 0
	numTableS := make([]int, 10)
	numTableG := make([]int, 10)
	for i := 0; i < len(secretChar); i++ {
		if secretChar[i] == guessChar[i] {
			Bull++
			Cow--
		}
		numTableG[int(guessChar[i])-int('0')]++
		numTableS[int(secretChar[i])-int('0')]++
	}
	for i := 0; i <= 9; i++ {
		if numTableS[i] < numTableG[i] {
			Cow += numTableS[i]
		} else {
			Cow += numTableG[i]
		}
	}
	return strconv.Itoa(Bull) + "A" + strconv.Itoa(Cow) + "B"
}

func gameOfLife(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			trans(board, i, j)
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 3 {
				board[i][j] = 0
			}
			if board[i][j] == 2 {
				board[i][j] = 1
			}
		}
	}
	fmt.Println(board)
}

func trans(board [][]int, I, J int) {
	count0, count1 := 0, 0
	for i := I - 1; i <= I+1; i++ {
		if i < 0 || i >= len(board) {
			continue
		}
		for j := J - 1; j <= J+1; j++ {
			if j < 0 || j >= len(board[I]) {
				continue
			}
			if i == I && j == J {
				continue
			}
			if board[i][j] == 0 || board[i][j] == 2 {
				count0++
			} else {
				count1++
			}
		}
	}
	if board[I][J] == 1 {
		if count1 < 2 || count1 > 3 {
			board[I][J] = 3
		}
	}
	if board[I][J] == 0 {
		if count1 == 3 {
			board[I][J] = 2
		}
	}
}

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	return true
}

func (this *Iterator) next() int {
	return 1
}

type PeekingIterator struct {
	Iter      *Iterator
	NextCache []int
}

//func Constructor(iter *Iterator) *PeekingIterator {
//	PeekIter:=PeekingIterator{
//		Iter:iter,
//		NextCache: make([]int,0),
//	}
//	return &PeekIter
//}

func (this *PeekingIterator) hasNext() bool {
	return len(this.NextCache) > 0 || this.Iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if len(this.NextCache) > 0 {
		nt := this.NextCache[0]
		this.NextCache = this.NextCache[1:]
		return nt
	} else {
		return this.Iter.next()
	}
}

func (this *PeekingIterator) peek() int {
	if len(this.NextCache) > 0 {
		return this.NextCache[0]
	} else {
		nt := this.Iter.next()
		this.NextCache = append(this.NextCache, nt)
		return nt
	}
}

func nthUglyNumber(n int) int {
	var (
		n2 = 0
		n3 = 0
		n5 = 0
		dp = make([]int, n+1)
	)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = min(dp[n2]*2, dp[n3]*3, dp[n5]*5)
		if dp[i] == dp[n2]*2 {
			n2++
		}
		if dp[i] == dp[n3]*3 {
			n3++
		}
		if dp[i] == dp[n5]*5 {
			n5++
		}
	}
	return dp[n-1]
}
func hIndex(citations []int) int {
	sort.Ints(citations)
	maxH := 0
	for i := len(citations) - 1; i >= 0; i-- {
		if len(citations)-i > maxH && citations[i] > maxH {
			maxH = min(len(citations)-i, citations[i])
		}
	}
	return maxH
}

func min(a ...int) int {
	if len(a) <= 0 {
		panic("ERROR")
	}
	m := a[0]
	for i := 0; i < len(a); i++ {
		if m > a[i] {
			m = a[i]
		}
	}
	return m
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
