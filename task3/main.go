package main

import (
	"fmt"
	"sync"
)

func square(k int) {
	sum += k * k
}


var sum int


func main() {
	a := [5]int{2, 4, 6, 8, 10}
	wag := sync.WaitGroup{}
	for _, val := range a {
		wag.Add(1)

		go func(val int) {
			square(val)
			wag.Done()
		}(val)
	}
	wag.Wait()
	fmt.Println("sum is - ", sum)
}
