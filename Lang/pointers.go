package main

import "fmt"

func changeValue1(str *string) {
	*str = "changed!"
}

func changeValue2(str string) {
	str = "changed!"
}

func main() {
	x := 7
	y := &x
	fmt.Println(x, y)
	*y = 8
	fmt.Println(x, y)

	toChange1 := "Hello"
	fmt.Println(toChange1)
	changeValue1(&toChange1)
	fmt.Println(toChange1)

	toChange2 := "Holla"
	fmt.Println(toChange2)
	changeValue2(toChange2)
	fmt.Println(toChange2)

	toChange3 := "hey"
	var pointer *string = &toChange3
	fmt.Println(pointer, *pointer, &pointer)

}
