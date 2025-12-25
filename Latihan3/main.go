package main

import "fmt"

func main() {
	ticket := map[int]string{
		1: "Ticket A",
		2: "Ticket B",
	}
	fmt.Println(ticket[2])
	verifyAvailability(ticket, 2)
	verifyAvailability(ticket, 3)
}
