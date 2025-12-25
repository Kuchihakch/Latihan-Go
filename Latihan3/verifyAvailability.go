package main

import "fmt"

// kurang bagus, ada case kalau ticket ada && value == "" -> anggapannya tidak ada
/* func verifyAvailability(ticket map[int]string, key int) {
	switch ticket[key] {
	case "":
		fmt.Println("Tiket Tidak Ditemukan")
	default:
		fmt.Println("Tiket Ditemukan")
	}
} */

//alternatif: (lebih valid)
func verifyAvailability(ticket map[int]string, key int) {
	_, ok := ticket[key] // value && bool
	fmt.Println(ok)
	if !ok {
		fmt.Println("Ticket Tidak Ditemukan")
	} else {
		fmt.Println("Tiket Ditemukan")
	}
}
