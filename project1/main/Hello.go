package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrintTree(root *TreeNode) {
	if root.Left != nil {
		PrintTree(root.Left)
	}
	if root.Right != nil {
		PrintTree(root.Right)
	}
}

func main() {
	head := Constructor()
	head.Insert("apple")
	fmt.Println(head.Search("apple"))
	fmt.Println(head.Search("app"))
	fmt.Println(head.StartsWith("app"))
	head.Insert("apple")
	fmt.Println(head.Search("apple"))
}

type Trie struct {
	Val         byte
	Next        map[byte]*Trie
	endWithThis bool
}

const NodeCount = 26 + 4

/** Initialize your data structure here. */
func Constructor() Trie {
	head := Trie{
		Val:         ' ',
		Next:        make(map[byte]*Trie),
		endWithThis: false}
	return head
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	head := this
	index := 0
	wlen := len(word)
	for index < wlen {
		w := word[index]
		if head.Next[w-'a'] != nil {
			head = head.Next[w-'a']
		} else {
			head.Next[w-'a'] = &Trie{
				Val:         w,
				Next:        make(map[byte]*Trie),
				endWithThis: false}
			head = head.Next[w-'a']
		}
		index++
	}
	head.endWithThis = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	head := this
	index := 0
	wlen := len(word)
	for index < wlen {
		w := word[index]
		if head.Next[w-'a'] != nil {
			head = head.Next[w-'a']
			index++
			continue
		}
		return false
	}
	if !head.endWithThis {
		return false
	}
	return true

}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	head := this
	index := 0
	wlen := len(prefix)
	for index < wlen {
		w := prefix[index]
		if head.Next[w-'a'] != nil {
			head = head.Next[w-'a']
			index++
			continue
		}
		return false
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
