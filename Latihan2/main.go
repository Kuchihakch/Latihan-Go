package main

import "fmt"

func main() {
	listAngka := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	Genap := filterEven(listAngka...)
	fmt.Println(Genap)
}
