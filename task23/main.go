package main

import "fmt"

func main() {
	a := []string{"A", "B", "C", "D", "E"}
	i := 2

	// Удалить элемент по индексу i из a.

	// 1. Копировать последний элемент в индекс i.
	a[i] = a[len(a)-1]

	// 2. Усечь срез.
	a = a[:len(a)-1]

	fmt.Println(a)
}
