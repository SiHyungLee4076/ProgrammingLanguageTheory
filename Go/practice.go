package main

import "fmt"

func main() {
	count := 0
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if count > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%d", i)
			count++
			if count == 5 {
				fmt.Println()
				count = 0
			}
		}
	}
	if count > 0 {
		fmt.Println()
	}
}
