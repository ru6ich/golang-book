package main

import (
	"fmt"
	"unicode"
)

func mergeSpace(s []byte) []byte {
	writeIdx := 0
	for i := 0; i < len(s); i++ {
		if unicode.IsSpace(rune(s[i])) {
			if writeIdx > 0 && s[writeIdx-1] != ' ' {
				s[writeIdx] = ' '
				writeIdx++
			}
		} else {
			s[writeIdx] = s[i]
			writeIdx++
		}
	}
	return s[:writeIdx]
}

func main() {
	s := []byte("Hello\t\t World   \t   \t\teveryone\t   !")
	fmt.Println(string(s))
	fmt.Println(string(mergeSpace(s)))
}
