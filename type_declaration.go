package main

import "fmt"

func main() {
	type NoKTP string
	var ktpEko NoKTP = "123123123"

	var contoh string = "222222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpEko)
	fmt.Println(contohKtp)
}
