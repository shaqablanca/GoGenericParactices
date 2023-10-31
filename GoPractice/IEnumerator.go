package main

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
