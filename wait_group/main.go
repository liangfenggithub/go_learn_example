package main

import (
	"fmt"
	"sync"
)

// To wait for multiple goroutines to finish, we can use a wait group.
var wg sync.WaitGroup

func main() {
	c := make(chan bool, 1)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(c, i)
	}
	// <-c
	wg.Wait()
	fmt.Println("all worker done")
}

func worker(c chan bool, index int) {
	defer wg.Done()
	sum := 0
	fmt.Println("worker", index, "begin work")
	for i := 0; i < 10000; i++ {
		sum += i
	}
	fmt.Println("worker ", index, "complete work, sum is ", sum)
	// c <- true
}
