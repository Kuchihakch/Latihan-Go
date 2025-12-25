package main

import "fmt"

func printEach(num ...int8) {
	for _, nums := range num {
		fmt.Print(nums, " ")
	}
	fmt.Println()
}
