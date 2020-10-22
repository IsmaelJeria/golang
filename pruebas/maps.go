package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m1["a"] = 10
	m1["b"] = 9

	fmt.Println(m1)
}
