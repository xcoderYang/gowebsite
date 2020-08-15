package main

import (
	"fmt"
	"time"
	"union_find/application"
)

/**

	并查集图

	0 ---- 1 ---- 2		3 ---- 4
	|	   |			|	   |
	|	   |			|	   |
	5 ---- 6 ---- 7		8	   9

**/

func main() {
	const n = 1000
	const count = 1000
	const wcount = 10
	const p = 0.592746
	var index2 = 0
	for index2 < wcount {
		var p_shark = 0
		var index = 0
		for index < count {
			time.Sleep(5 * time.Millisecond)
			if quick_union_test(n, p) {
				p_shark++
			}
			index++
		}
		fmt.Println(float64(p_shark) / float64(count))
		index2++
	}

	// var qf = quick_union_improve_compress_init(10)
	// qf.union(4, 3)
	// qf.union(3, 8)
	// qf.union(6, 5)
	// qf.union(9, 4)
	// qf.union(2, 1)
	// qf.union(5, 0)

	// qf.union(7, 2)
	// qf.union(6, 1)
	// qf.union(7, 3)

	// fmt.Println(qf.connected(0, 7))
	// fmt.Println(qf.connected(8, 9))
	// fmt.Println(qf.connected(5, 4))
	// fmt.Println(qf.id)
}

func quick_union_test(n int, p float64) bool {
	var pers = application.Percolation_init(n, p)
	var matrix = pers.GetMatrix()
	var qu = quick_union_improve_compress_application_init(n)

	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			// fmt.Println(i, j)
			if matrix[i][j] == 0 {
				if i == 0 {
					qu.union(j, 0)
				}
				if i == n-1 {
					qu.union(i*n+j, n*n-1)
				}
				if i > 0 && matrix[i-1][j] == 0 {
					qu.union(i*n+j, (i-1)*n+j)
				}
				if i < n-1 && matrix[i+1][j] == 0 {
					qu.union(i*n+j, (i+1)*n+j)
				}
				if j > 0 && matrix[i][j-1] == 0 {
					qu.union(i*n+j, i*n+j-1)
				}
				if j < n-1 && matrix[i][j+1] == 0 {
					qu.union(i*n+j, i*n+j+1)
				}
			}
		}
	}
	// fmt.Println(qu.connected(0, n*n-1))
	return qu.connected(0, n*n-1)
	// i := 0
	// j := 0
	// for i < n {
	// 	j = 0
	// 	for j < n {
	// 		fmt.Print(qu.id[i*n+j], " ")
	// 		j++
	// 	}
	// 	fmt.Println("")
	// 	i++
	// }
}

type quick_union_improve_compress_application struct {
	id []int
	sz []int
}

func quick_union_improve_compress_application_init(n int) quick_union_improve_compress_application {
	index := 0
	var qu quick_union_improve_compress_application
	qu.id = make([]int, n*n)
	qu.sz = make([]int, n*n)
	for index < n*n {
		qu.id[index] = index
		qu.sz[index] = 1
		index++
	}
	return qu
}
func (qu quick_union_improve_compress_application) root(n int) int {
	for n != qu.id[n] {
		n = qu.id[n]
		qu.id[n] = qu.id[qu.id[n]]
	}
	return n
}
func (qu quick_union_improve_compress_application) connected(p int, q int) bool {
	return qu.root(q) == qu.root(p)
}
func (qu quick_union_improve_compress_application) union(p int, q int) {
	pRoot := qu.root(p)
	qRoot := qu.root(q)
	sz := qu.sz
	id := qu.id
	if pRoot != qRoot {
		if sz[pRoot] > sz[qRoot] {
			id[qRoot] = id[pRoot]
			sz[pRoot] += sz[qRoot]
		} else {
			id[pRoot] = id[qRoot]
			sz[qRoot] += sz[pRoot]
		}
	}
}

/**

	quick_find 算法

	如果这些节点是连通的 <=> 每个节点对应的值都是一样的

	方法时间复杂度:
		初始化: O(n)
		union: O(n)
		find:  O(1)

	构建 union操作总耗时 O(n^2) 太慢

	还有个不对称问题，对于边的输入有顺序的要求

**/
func quick_find_init(n int) quick_find {
	var qf quick_find
	qf.id = make([]int, n)
	for i, _ := range qf.id {
		qf.id[i] = i
	}
	return qf
}

type quick_find struct {
	id []int
}

func (qf quick_find) union(p int, q int) {
	id := qf.id
	pid := id[p]
	qid := id[q]
	for i, _ := range id {
		if id[i] == pid {
			id[i] = qid
		}
	}
}
func (qf quick_find) connected(p int, q int) bool {
	return qf.id[p] == qf.id[q]
}

/**

	quick_union 算法

	如果这些节点是连通的 <=> 每个节点的根是一样的

	方法时间复杂度: 最差情况，整颗树一个分支，相当于一个序列
		初始化: O(n)
		union: O(n)
		find:  O(n)

	构建 union操作总耗时 O(n^2) 太慢

**/

func quick_union_init(n int) quick_union {
	var qf quick_union
	qf.id = make([]int, n)
	for i, _ := range qf.id {
		qf.id[i] = i
	}
	return qf
}

type quick_union struct {
	id []int
}

func (qf quick_union) root(p int) int {
	id := qf.id
	if id[p] != p {
		p = id[p]
	}
	return p
}

func (qf quick_union) union(p int, q int) {
	pRoot := qf.root(p)
	qRoot := qf.root(q)
	qf.id[qRoot] = pRoot
}
func (qf quick_union) connected(p int, q int) bool {
	return qf.root(p) == qf.root(q)
}

/**

	quick_union_improve算法

	在 quick_union算法上增加 sz，用来比较每个root节点上的负载
	保证整棵树足够的扁平

**/
func quick_union_improve_init(n int) quick_union_improve {
	var qf quick_union_improve
	qf.id = make([]int, n)
	qf.sz = make([]int, n)
	for i, _ := range qf.id {
		qf.id[i] = i
		qf.sz[i] = 1
	}
	return qf
}

type quick_union_improve struct {
	id []int
	sz []int
}

func (qf quick_union_improve) root(p int) int {
	id := qf.id
	if id[p] != p {
		p = id[p]
	}
	return p
}

func (qf quick_union_improve) union(p int, q int) {
	id := qf.id
	sz := qf.sz
	pRoot := qf.root(p)
	qRoot := qf.root(q)
	if qf.sz[pRoot] > qf.sz[qRoot] {
		id[qRoot] = id[pRoot]
		sz[pRoot] += sz[qRoot]
	} else {
		id[pRoot] = id[qRoot]
		sz[qRoot] += sz[pRoot]
	}
}
func (qf quick_union_improve) connected(p int, q int) bool {
	return qf.root(p) == qf.root(q)
}

/**




**/

func quick_union_improve_compress_init(n int) quick_union_improve_compress {
	var qf quick_union_improve_compress
	qf.id = make([]int, n)
	qf.sz = make([]int, n)
	for i, _ := range qf.id {
		qf.id[i] = i
		qf.sz[i] = 1
	}
	return qf
}

type quick_union_improve_compress struct {
	id []int
	sz []int
}

func (qf quick_union_improve_compress) root(p int) int {
	id := qf.id
	if id[p] != p {
		p = id[p]
		id[p] = id[id[p]]
	}
	return p
}

func (qf quick_union_improve_compress) union(p int, q int) {
	id := qf.id
	sz := qf.sz
	pRoot := qf.root(p)
	qRoot := qf.root(q)
	if qf.sz[pRoot] > qf.sz[qRoot] {
		id[qRoot] = id[pRoot]
		sz[pRoot] += sz[qRoot]
	} else {
		id[pRoot] = id[qRoot]
		sz[qRoot] += sz[pRoot]
	}
}
func (qf quick_union_improve_compress) connected(p int, q int) bool {
	return qf.root(p) == qf.root(q)
}
