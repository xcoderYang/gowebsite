package main

func main() {

}

type stack struct {
	cache []int
}

func (sk stack) push(n int) {
	sk.cache = append(sk.cache, n)
}

func (sk stack) length() int {
	return len(sk.cache)
}
func (sk stack) pop() int {
	if sk.length() == 0 {
		return 0
	}
	item := sk.cache[sk.length()-1]
	sk.cache = sk.cache[:len(sk.cache)-1]
	return item
}
func (sk stack) isEmpty() {

}
