package main

import "fmt"

func main() {
	num := []int{6, 5, 3, 7, 8, 9, 2, 1, 7, 8, 10, 2}
	quickSort(num)
	fmt.Println(num)
}

func quickSort(arr []int) []int {
	sort(arr, 0, len(arr)-1)
	return arr
}

func sort(arr []int, low int, high int) {
	if high <= low {
		return
	}
	j := partition(arr, low, high)
	sort(arr, low, j-1)
	sort(arr, j+1, high)
}

func sort3way(arr []int, low int, high int) {
	if high <= low {
		return
	}
	lt, gt, v := low, high, arr[low]
	i := low
	for i <= gt {
		cmp := arr[i] - v
		if cmp < 0 {
			exch(arr, lt, i)
			lt++
			i++
		} else if cmp > 0 {
			exch(arr, i, gt)
			gt--
		} else {
			i++
		}
	}
	sort3way(arr, low, lt-1)
	sort3way(arr, gt+1, high)
}

func partition(arr []int, low int, high int) int {
	i, j := low+1, high
	for true {
		for arr[i] < arr[low] {
			i++
			if i == high {
				break
			}
		}
		for arr[low] < arr[j] {
			j--
			if j == low {
				break
			}
		}
		if i >= j {
			break
		}
		exch(arr, i, j)
	}
	exch(arr, low, j)
	return j
}
func exch(arr []int, a int, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}

func medianOf3(arr []int, low int, high int) int {
	mid := (high + low) / 2
	if arr[high] > arr[mid] {
		if arr[mid] > arr[low] {
			return mid
		} else if arr[low] < arr[high] {
			return low
		} else {
			return high
		}
	} else {
		if arr[mid] < arr[low] {
			return mid
		} else if arr[low] < arr[high] {
			return high
		} else {
			return low
		}
	}
}
