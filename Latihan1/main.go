package main

import "fmt"

func main() {
	Numbers := []int8{1, 2, 3, 4, 5}
	fmt.Println(Numbers)
	//cek sumAll
	Sum := SumAll(Numbers...)
	fmt.Println(Sum)
	//cek printEach
	printEach(Numbers...)
	var appendSlice = append(Numbers, 6, 7)
	printEach(appendSlice...) // 6,7 ada di data baru
	printEach(Numbers...)     // harusnya tetap sama, karena capnya penuh yaitu 5 tanpa append
}
