package main

import "fmt"

type Student struct {
	name  string
	grade []int
	age   int
}

func (s1 Student) getAge() int {
	return s1.age
}

func (s2 *Student) setAge(age int) {
	s2.age = age
}

func (s3 *Student) getAverageGrade() float32 {
	sum := 0

	for _, v := range s3.grade {
		sum += v
	}

	return float32(sum) / float32(len(s3.grade))
}

func (s4 *Student) getMaxGrade() int {
	curMax := 0

	for _, v := range s4.grade {
		if curMax < v {
			curMax = v
		}
	}
	return curMax
}

func main() {
	s1 := Student{"Tim", []int{23, 45, 67}, 12}
	fmt.Println(s1.getAge())

	s2 := Student{"Tim", []int{23, 45, 67}, 50}
	fmt.Println(s2)
	s2.setAge(60)
	fmt.Println(s2)

	s3 := Student{"Jim", []int{23, 43, 45}, 70}
	average := s3.getAverageGrade()
	fmt.Println(average)

	s4 := Student{"Jim", []int{55, 33, 22, 43, 45}, 70}
	max1 := s4.getMaxGrade()
	max2 := s2.getMaxGrade()
	fmt.Println(max1)
	fmt.Println(max2)
	fmt.Println(max1, max2)
}
