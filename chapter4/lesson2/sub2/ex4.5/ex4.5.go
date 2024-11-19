package main

import "fmt"

func removeDuplicate(slice []string) []string {
	idx := 1
	for i := 1; i < len(slice); i++ {
		if slice[i] != slice[i-1] {
			slice[idx] = slice[i]
			idx++

		}
	}
	return slice[:idx]
}

func main() {
	s := []string{"apple", "apple", "banana", "banana", "apple", "orange", "orange", "orange"}
	fmt.Println(s)
	fmt.Println(removeDuplicate(s))
}
