package main

import "fmt"

// func main() {
// 	const firstName string = "Brama"
// 	const lastName = "Tri Hadi"

// 	// error
// 	// firstName = "Budi"
// 	// lastName = "Joko"
// }

func main() {
	const (
		firstName string = "Brama"
		lastName         = "Tri Hadi"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
