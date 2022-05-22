package main

import (
	"fmt"

	"os"

	"github.com/aleksandar-had/gophercises/linksparser"
)

func main() {
	r, err := os.Open("ex1.html")
	//	r := strings.NewReader(exampleHtml)
	if err != nil {
		panic(err)
	}

	links, err := linksparser.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
