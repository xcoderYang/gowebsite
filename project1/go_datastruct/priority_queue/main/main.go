package main

import "fmt"

func main() {
	datas := []int{
		4, 1, 5, 7, 8, 1, 2, 12, 63, 64, 23, 46, 76, 7, 123, 12535, 3,
	}
	datas = buildHead(datas)
	//fmt.Println(datas)
}

func less(a int, b int) bool {
	return a < b
}

func exch(pa *int, pb *int) {
	temp := *pa
	*pa = *pb
	*pb = temp
}

func swim(k int, datas *[]int) {
	for k > 1 && less((*datas)[k/2], (*datas)[k]) {
		exch(&(*datas)[k/2], &(*datas)[k])
		k = k / 2
	}
}

func insert(x int, datas []int) []int {
	datas = append(datas, x)
	//fmt.Println(datas)
	swim(len(datas)-1, &datas)
	return datas
}

func sink(k int, datas []int) {
	N := len(datas) - 2
	for 2*k <= N {
		j := 2 * k
		// 选取 j和 j+1中的最大值,假设最大值为 j，如果 j小于 j+1，则 j++变成 j+1
		// j这个变量永远指向子元素中的最大值
		if j < N && less(datas[j], datas[j+1]) {
			j++
		}
		// 如果父节点比子节点要大，说明不用下沉了
		if !less(datas[k], datas[j]) {
			break
		}
		// 下沉
		exch(&datas[k], &datas[j])
		k = j
	}
}

func delMax(datas *[]int) int {
	max := (*datas)[1]
	exch(&(*datas)[1], &(*datas)[len((*datas))-1])
	sink(1, (*datas))
	*datas = (*datas)[:len((*datas))-1]
	return max
}

func buildHead(datas []int) []int {
	ans := make([]int, 1)
	for i := 0; i < len(datas); i++ {
		ans = insert(datas[i], ans)
	}
	cache := len(ans)
	for i := 0; i < cache-1; i++ {
		fmt.Println(delMax(&ans))
	}
	return ans
}
