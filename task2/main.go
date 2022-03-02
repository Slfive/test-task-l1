package main

import (
	"fmt"
	"sync"
)

func square(c int) {
	fmt.Println(c * c)
}

func main() {
	a := [4]int{3, 4, 6, 7}
	wag := sync.WaitGroup{}
	for _, v := range a {
		wag.Add(1)
		go func() {
			square(v)
			wag.Done()
		}()
		wag.Wait()
	}
	fmt.Println("Done")

}