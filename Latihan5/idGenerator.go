package main

func idGenerator(n int) []int {
	ids := []int{}
	for i := 1; i <= n; i++ {
		ids = append(ids, i)
	}
	return ids
}
