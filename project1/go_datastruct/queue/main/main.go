package main

import (
	"errors"
	"fmt"
)

type queue struct {
	First *node
	Last  *node
	Len   int
}

type node struct {
	Val  interface{}
	Next *node
	Pre  *node
}

func main() {
	var que queue
	que.enqueue("123")
	que.enqueue("456")
	que.enqueue("789")
	que.enqueue("101")
	fmt.Println(que.dequeue())
	fmt.Println(que.getLength())
	fmt.Println(que.dequeue())
	fmt.Println(que.getLength())
	fmt.Println(que.dequeue())
	fmt.Println(que.getLength())
	fmt.Println(que.dequeue())
	fmt.Println(que.isEmpty())
	fmt.Println(que.getLength())
}

func (qu *queue) enqueue(data interface{}) {
	nNode := &node{
		Val:  data,
		Pre:  qu.First,
		Next: nil}
	if qu.First == nil {
		qu.First = nNode
	} else {
		qu.First.Next = nNode
		qu.First = nNode
	}
	if qu.Last == nil {
		qu.Last = nNode
	}
	qu.Len++
}

func (qu *queue) dequeue() interface{} {
	if qu.Len > 0 {
		nNode := qu.Last.Val
		if qu.Last.Next != nil {
			qu.Last.Next.Pre = nil
		}
		qu.Last = qu.Last.Next
		qu.Len--
		return nNode
	}
	return errors.New("error")
}

func (qu queue) isEmpty() bool {
	return qu.Len <= 0
}

func (qu queue) getLength() int {
	return qu.Len
}
