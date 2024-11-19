package main

import "fmt"

func rotate1(s1 []int) []int {
	for i := 0; i < len(s1)/2; i++ {
		s1[i], s1[len(s1)-1-i] = s1[len(s1)-1-i], s1[i]
	}
	return s1
}

func rotate2(s1 []int) {
	for i := 0; i < len(s1)/2; i++ {
		s1[i], s1[len(s1)-1-i] = s1[len(s1)-1-i], s1[i]
	}
}

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1)
	s2 := rotate1(s1)
	fmt.Println(s2)

	s3 := []int{1, 2, 3, 4, 5}
	fmt.Println(s3)
	rotate2(s3)
	fmt.Println(s3)
}
