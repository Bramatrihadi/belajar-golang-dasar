package main

import "fmt"

func main() {
	names := [...]string{
		"Brama",
		"Tri",
		"Hadi",
		"Syapoetra",
		"Eko",
		"Kurniawan",
	}
	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	days := [...]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}

	daySlice1 := days[5:] // sabtu, minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "sabtu lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Brama"
	newSlice[1] = "Brama"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Brama")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
	//di golang lebih sering menggunakan slice
}
