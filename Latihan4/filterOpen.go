package main

func filterOpen(status ...string) int {
	count := 0
	for _, s := range status {
		if s == "open" {
			count++
		}
	}
	return count
}
