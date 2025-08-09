package main

import (
	"fmt"
)

func main() {
	var num1 int = 9
	var num2 int = 4
	answer := num1 / num2
	fmt.Printf("%g", answer)

	var num3 float32 = 9
	var num4 float32 = 4
	answer1 := num3 / num4
	fmt.Printf("%g", answer1)

	var num6 int = 9
	var num5 int = 4
	answer2 := num6 % num5
	fmt.Printf("%d", answer2)

	x := 5
	y := 6
	val := x < y
	z := 4.9
	val1 := float64(z) > 5
	fmt.Printf("%t", val)
	fmt.Printf("%t", val1)

	a := "Tim"
	b := "tim"
	com := a == b
	fmt.Printf("%t", com)

	pare := true || false && true
	fmt.Printf("%t", pare)

	pare1 := (true || false) && !false
	pare2 := pare1 == com
	fmt.Printf("%t\n", pare1)
	fmt.Printf("%t\n", pare2)

	nick := "time"
	fmt.Printf("Before if\n")
	if nick == "tim" {
		fmt.Printf("Inside if\n")
	}
	fmt.Printf("After if\n")

	age := 10
	if age >= 18 {
		fmt.Println(("You can ride"))
	} else {
		fmt.Println(("You can't ride"))
		fmt.Printf("Wait %d years\n", 18-age)
	}

	age1 := 10
	if age1 >= 18 {
		fmt.Println("You can ride")
	} else if age1 >= 14 {
		fmt.Println("You can ride with a parent")
	} else {
		fmt.Println("You can't ride")
	}

	g := 0
	for g <= 5 {
		fmt.Println(g)
		g++
	}

	for h := 10; h <= 15; h += 2 {
		fmt.Println(h)
	}

	for i := 10; i <= 1000; i += 2 {
		if i != 0 && i%3 == 0 && i%7 == 0 && i%9 == 0 {
			fmt.Println(i)
			continue
		}
	}

	for k := 10; k <= 1000; k += 2 {
		if k != 0 && k%3 == 0 && k%7 == 0 && k%9 == 0 {
			fmt.Println(k)
			break
		}
	}

	ans := -1

	switch ans {
	case 1, -1:
		fmt.Println("1. one")
		fmt.Println("2. two")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not a case")
	}

	wer := -1
	switch {
	case wer > 0:
		fmt.Println("Greater than zero")
	case wer < 0:
		fmt.Println("Less than zero")
	default:
		fmt.Println("zero")

	}

	var arr [5]string
	fmt.Println(arr)

	var brr [5]int
	brr[0] = 100
	brr[4] = 20
	fmt.Println(brr)

	crr := [3]int{4, 5, 6}
	fmt.Println(crr)
	fmt.Println(len(crr))

	drr := []int{10, 11, 12}
	sum := 0

	for k := 0; k < len(drr); k++ {
		sum += drr[k]
	}

	fmt.Println(sum)

	arr2D := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(arr2D)
	fmt.Println(arr2D[1][1])

	var err [5]int = [5]int{1, 2, 3, 4, 5}
	var frr []int = err[1:3]
	fmt.Println(frr)
	fmt.Println(len(frr))
	fmt.Println(cap(frr))
	fmt.Println(frr[:cap(frr)])

	var grr []int = []int{5, 6, 7, 8, 9}
	fmt.Println(cap(grr))
	fmt.Println(cap(grr[:3]))
	hrr := append(grr, 10)
	fmt.Println(hrr)

	irr := make([]int, 5)
	fmt.Println(irr)
	fmt.Printf("%T", irr)

	var jrr []int = []int{1, 2, 3, 4, 9, 5, 6, 7, 8, 7, 9}
	for i := 0; i < len(jrr); i++ {
		fmt.Println(jrr[i])
	}

	for _, element := range jrr {
		fmt.Printf("%d\n\n", element)
	}

	for i, element := range jrr {
		for j, element2 := range jrr {
			if element == element2 && j > i {
				fmt.Println(element)
			}
		}
	}

	var krr []int = []int{1, 2, 3, 4, 5, 6, 5, 7, 8, 9}
	for o, elem1 := range krr {
		for p := o + 1; p < len(krr); p++ {
			elem2 := krr[p]
			if elem1 == elem2 {
				fmt.Println(elem1)
			}

		}
	}

	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 7,
	}
	fmt.Println(mp)

	mp1 := map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 7,
	}
	fmt.Println(mp1)

	mp2 := make(map[string]int)
	fmt.Println(mp2)

	fmt.Println(mp["apple"])
	mp["apple"] = 900
	mp["pawpaw"] = 700
	fmt.Println(mp)

	delete(mp, "apple")
	fmt.Println(mp)

	value, ok := mp["tim"]
	fmt.Println(value, ok)
	fmt.Println(mp)

	value1, ok1 := mp["apple"]
	fmt.Println(value1, ok1)

	value2, ok2 := mp["pear"]
	fmt.Println(value2, ok2)
	fmt.Println(len(mp))

}
