package main

import "fmt"

type Point struct {
	x int32
	y int32
	z bool
}

type Circle1 struct {
	radius float64
	center *Point
}

type Circle2 struct {
	diameter float64
	*Point
}

func changeX(a *Point) {
	a.x = 100
}

func main() {
	var p1 Point = Point{1, 2, false}
	var p2 Point = Point{3, 4, true}
	fmt.Println(p1.x, p1.y, p1.z)
	fmt.Println(p2.x, p2.y, p2.z)

	var p3 Point = Point{y: 3}
	fmt.Println(p3)

	p4 := &Point{z: true}
	fmt.Println(p4)

	p5 := &Point{y: 7}
	fmt.Println(p5)
	changeX(p5)
	fmt.Println(p5)

	p6 := Point{10, 20, true}
	c1 := Circle1{5.5, &p6}
	c2 := Circle1{9.9, &Point{10, 20, false}}
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c1.center.x)
	fmt.Println(c1.center.y)
	fmt.Println(c1.center.z)

	c3 := Circle2{3.0, &p6}
	c4 := Circle2{5.6, &Point{11, 22, false}}
	fmt.Println(c3)
	fmt.Println(c4.x)
	fmt.Println(c4.y)
	fmt.Println(c4.z)
}
