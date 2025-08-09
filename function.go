package main

import "fmt"

func test(x int) {
	fmt.Println(x)
}
func add(x, y int) int {
	return x + y
}
func add_2(x, y int) (int, int) {
	return x + y, x * y
}
func add_3(x, y int) (z1, z2 int) {
	defer fmt.Println("hello")
	z1 = x + y
	z2 = x - y
	defer fmt.Println("before return")
	return
}
func j_test(jt int) {
	fmt.Println("Just another test", jt)
}
func b_test(myFunc func(x int) int) {
	fmt.Println(myFunc(7))
}
func returnFunc(x string) func() {
	return func() { fmt.Println(x) }
}

func main() {
	test(5)
	test(9)
	ans := add(13, 7)
	fmt.Println(ans)
	ans1, ans2 := add_2(20, 30)
	fmt.Println(ans1, ans2)
	ans3, ans4 := add_3(30, 40)
	fmt.Println(ans3, ans4)
	jt := j_test
	jt(5)

	test := func(d int) {
		fmt.Println("Just a hello", d)
	}
	test(4)

	a_test := func(x int) int {
		return x * -1
	}
	fmt.Println(a_test(8))

	c_test := func(y int) int {
		return y + 2
	}
	b_test(c_test)

	func() {
		fmt.Println("Hello test")
	}()

	returnFunc("hello")()
	xx := returnFunc(("Goodbye"))
	xx()

}
