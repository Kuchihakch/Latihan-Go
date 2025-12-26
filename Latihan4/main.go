package main

import "fmt"

func main() {
	ListStatus := []string{"open", "closed", "open", "open"}
	openStatus := filterOpen(ListStatus...)
	fmt.Println("Jumlah Open: ", openStatus)
}
