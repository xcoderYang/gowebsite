package  main

import (
	"fmt"
)

type ListNode struct{
	Val int
	Next *ListNode
}

func main(){
	// var case1 = ListNode{4,nil}
	// *case1.Next = ListNode{2,nil}
	// *case1.Next.Next = ListNode{1,nil}
	// *case1.Next.Next.Next = ListNode{3,nil}
	// fmt.Println(sortList(&case1))
	sortList(&ListNode{1, nil})
}
func sortList(head *ListNode) *ListNode {
	var merge func(head1 *ListNode, head2 *ListNode) *ListNode
	merge = func (head1 *ListNode, head2 *ListNode) *ListNode {
		var point1, point2 = head1, head2
		var prePoint *ListNode
		var headMain *ListNode
		if head1.Val <= head2.Val{
			headMain = head1
			prePoint = head1
			point1 = point1.Next
		}else{
			headMain = head2
			prePoint = head2
			point2 = point2.Next
		}
		for point1 != nil && point2 != nil{
			if point1.Val < point2.Val{
				prePoint.Next = point1
				point1 = point1.Next
			}else{
				prePoint.Next = point2
				point2 = point2.Next
			}
			prePoint = prePoint.Next
		}
		if point1 == nil{
			prePoint.Next = point2
		}else{
			prePoint.Next = point1
		}
		return headMain
	}
	
	var mergeSort func(head *ListNode) *ListNode
	mergeSort = func(head *ListNode) *ListNode{
		if head.Next == nil{
			return head
		}
		var quick, slow = head, head
		var preslow = head
		for quick != nil{
			slow = slow.Next
			quick = quick.Next
			if slow != head{
				preslow = preslow.Next
			}
			if quick != nil{
				quick = quick.Next
			}
		}
		preslow.Next = nil
		var left = mergeSort(head)
		var right = mergeSort(slow)
		return merge(left, right)
	}

	// var test1 = ListNode{1, &ListNode{3, &ListNode{5, &ListNode{7, nil}}}}
	// var test2 = ListNode{2, &ListNode{4, &ListNode{6, &ListNode{8, nil}}}}
	// var ans = merge(&test1, &test2)
	var ans = mergeSort(&ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}})
	for ans!=nil{
		fmt.Println(ans.Val)
		ans = ans.Next
	}
	return &ListNode{0, nil}
}