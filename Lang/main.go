package main

import "fmt"

func main() {

	var name string
	name = "Hello Tim"

	// var number uint16 = 260
	// number = number + 5

	var number = 10
	var fl = 299.2
	var bl bool
	flo := 3823.3
	boo := true

	fmt.Println("Hello, Go", name, number, boo, bl)
	fmt.Printf("%T", number)
	fmt.Printf("%T", fl)
	fmt.Printf("%T", flo)
	fmt.Printf("%T %v", 500, "hi")
	fmt.Printf("%T %v", "hi", 500)
	var x string = fmt.Sprintf("%T %v", "hi", 500)
	fmt.Println(x)
	fmt.Printf("Hello %t", 10)
	fmt.Printf("Hello %t", false)
	fmt.Printf("Number: %b", 1025)
	fmt.Printf("Number: %o", 1025)
	fmt.Printf("Number: %d", 1025)
	fmt.Printf("Number: %x", 3425)
	fmt.Printf("Number: %X", 3425)
	fmt.Printf("Number: %e", 3.22352363463463)
	fmt.Printf("Number: %f", 3.22352363463463)
	fmt.Printf("Number: %g", 3.22352363463463)

	fmt.Printf("Name: %s", "tim")
	fmt.Printf("Name: %q", "tim")
	fmt.Printf("Name: %9s", "time")
	fmt.Printf("Name: %9q", "time")
	fmt.Printf("Name: %-9s is cool", "time")
	fmt.Printf("Name: %-9q is cool", "time")
	fmt.Printf("Number %07d", 39)

	var out string = fmt.Sprintf("Name: %-9q is not cool", "time")
	fmt.Println(out)
	var outer string = fmt.Sprintf("Name:\n %-9q is not cool", "time")
	fmt.Println(outer)
	var outer1 string = fmt.Sprintf("Name:\f %-9q is not cool", "time")
	fmt.Println(outer1)

}
