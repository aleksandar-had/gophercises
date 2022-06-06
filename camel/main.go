package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	fmt.Println("Input:", input)

	answer := 1
	if len(input) <= 1 {
		answer = len(input)
		fmt.Println(answer)
	}

	for _, ch := range input {
		// min, max := 'A', 'Z'
		// if ch >= min && ch <= max {
		// 	// A new capital letter found!
		// 	answer++
		// }

		str := string(ch)
		if strings.ToUpper(str) == str {
			answer++
		}
	}
	fmt.Println(answer)
}
