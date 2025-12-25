package main

func filterEven(num ...int) []int {
	even := []int{}
	for _, i := range num {
		if i%2 != 0 {
			continue
		}
		even = append(even, i)
	}
	return even
}
