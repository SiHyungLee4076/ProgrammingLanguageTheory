package main

import "fmt"

func main() {
	rawLiteral := `good morning\n have a good day\n`
	var interLiteral string = "good morning\n have a good day\n"
	fmt.Println(rawLiteral)
	fmt.Println(interLiteral)
}