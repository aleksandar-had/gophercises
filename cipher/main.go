package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var delta int
	var line string
	var ret []rune

	fmt.Scanf("%d\n", &delta)
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		line = input.Text()
	}

	if len(strings.TrimSpace(line)) == 0 {
		line = "Print-Out a single ZzZzZ, alright?"
	}

	for _, ch := range line {
		ret = append(ret, encodeCeaser(ch, delta))
	}
	fmt.Println(string(ret))
}

func encodeCeaser(r rune, delta int) rune {
	if r >= 'A' && r <= 'Z' {
		return rotate(r, 'A', delta)
	} else if r >= 'a' && r <= 'z' {
		return rotate(r, 'a', delta)
	}
	return r
}

func rotate(r rune, base int, delta int) rune {
	tmp := int(r) - base
	tmp = (tmp + delta) % 26

	return rune(tmp + base)
}
