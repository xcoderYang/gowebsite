package main

import (
	"sync"
	_ "time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var test1 = 12345.87654321
		var test2 = 12345.87654321
		for {
			test1 = 12345.87654321
			test1 = test1 * test2
		}
		wg.Done()
	}()
	go func() {
		var test1 = 12345.87654321
		var test2 = 12345.87654321
		for {
			test1 = 12345.87654321
			test1 = test1 * test2
		}
		wg.Done()
	}()
	go func() {
		var test1 = 12345.87654321
		var test2 = 12345.87654321
		for {
			test1 = 12345.87654321
			test1 = test1 * test2
		}
		wg.Done()
	}()
	wg.Wait()

}
