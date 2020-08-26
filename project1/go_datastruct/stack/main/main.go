package main

import "fmt"

func main() {

	// 栈的应用，执行 ( 1 + ( ( 2 + 3 ) * ( 4 * 5 ) ) )
	// fmt.Println(stackCompute("1+(((2+3)*(4*5))/5)-5)"))
	// a := (1 + ((2+3)*(4*5))/10)
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	testArr(a)
	fmt.Println(a)
}

func testArr(arr []int) {
	arr[0] = 10
}

func stackCompute(str string) int {
	var (
		vsk  stackLink
		opsk stackLink
	)
	for _, v := range str {
		if v <= '9' && v >= '0' {
			vsk.push(int(v) - '0')
		} else if v == '+' || v == '-' || v == '*' || v == '/' {
			opsk.push(int(v))
		} else if v == ')' {
			n1 := vsk.pop().(int)
			n2 := vsk.pop().(int)
			op := opsk.pop().(int)
			var ans int
			switch op {
			case '+':
				ans = n2 + n1
			case '-':
				ans = n2 - n1
			case '*':
				ans = n2 * n1
			case '/':
				ans = n2 / n1
			}
			vsk.push(int(ans))
		}
	}
	for !opsk.isEmpty() {
		n1 := vsk.pop().(int)
		n2 := vsk.pop().(int)
		op := opsk.pop().(int)
		var ans int
		switch op {
		case '+':
			ans = n2 + n1
		case '-':
			ans = n2 - n1
		case '*':
			ans = n2 * n1
		case '/':
			ans = n2 / n1
		}
		vsk.push(int(ans))
	}
	char := vsk.pop().(int)
	return int(char)
}

/**

	用数组实现栈

**/
type stack struct {
	cache []int
}

func (sk *stack) push(n int) {
	sk.cache = append(sk.cache, n)
}

func (sk *stack) length() int {
	return len(sk.cache)
}
func (sk *stack) pop() int {
	if sk.length() == 0 {
		return 0
	}
	item := sk.cache[sk.length()-1]
	sk.cache = sk.cache[:len(sk.cache)-1]
	return item
}
func (sk *stack) isEmpty() bool {
	return len(sk.cache) == 0
}

/**

	用链表实现栈

**/

type stackLink struct {
	Top    *node
	Length int
}

type node struct {
	Val  interface{}
	Prev *node
}

func (sl *stackLink) push(value interface{}) {
	newNode := &node{
		Val:  value,
		Prev: sl.Top}
	sl.Top = newNode
	sl.Length++
}

func (sl *stackLink) pop() interface{} {
	topNodeVal := sl.Top.Val
	sl.Top = sl.Top.Prev
	sl.Length--
	return topNodeVal
}

func (sl stackLink) length() int {
	return sl.Length
}

func (sl stackLink) isEmpty() bool {
	return sl.Length == 0
}

func (sl stackLink) peer() interface{} {
	return sl.Top.Val
}
