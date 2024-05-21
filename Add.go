package main

import (
	"fmt"
)

func add(x, y float32) float32 {
	return x + y
}
func swap(first, second string) (string, string) {
	return second, first
}
func add_final() {
	fmt.Println(add(20.1, 30))
	a, b := swap("HAi", "Hello")
	fmt.Println(a, b)
}
