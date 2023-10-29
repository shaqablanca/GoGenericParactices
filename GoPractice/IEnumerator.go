package main

import "fmt"

func createIterator(collection []int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for _, item := range collection {
			ch <- item
		}
	}()

	return ch
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	iterator := createIterator(data)

	for item := range iterator {
		fmt.Println(item)
	}
}
