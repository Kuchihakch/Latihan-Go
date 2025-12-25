package main

func SumAll(numbers ...int8) int {
	total := 0
	//for - range
	/* for _, number := range numbers {
		total += int(number)
	}
	return total */
	//manual
	for i := 0; i < len(numbers); i++ {
		total += int(numbers[i])
	}
	return total
}
