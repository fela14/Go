package main

import "fmt"

func changefirst(slice []int) {
	slice[0] = 1000
}

func main() {
	// Example 1: primitive values
	var a int = 5
	b := a
	b = 7
	fmt.Println(a, b) // 5 7  (a is unchanged)

	// Example 2: slices (reference type)
	x := []int{3, 4, 5}
	y := x
	y[0] = 6
	fmt.Println(x, y) // [6 4 5] [6 4 5]  (both changed)

	m := map[string]int{"hello": 3}
	n := m
	n["hi"] = 4
	fmt.Println(n, m)

	g := [3]int{10, 20}
	h := g
	h[0] = 50
	fmt.Println(g, h)

	s := []int{11, 12, 13}
	fmt.Println(s)
	changefirst(s)
	fmt.Println(s)
}
