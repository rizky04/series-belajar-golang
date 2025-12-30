package main

import "fmt"

func main() {
	//const adalah variable yang nilainya tidak bisa diubah ketika sudah diberi nilai
	// const name string = "amin"
	// const lastName = "jaya"

	// fmt.Println(name, lastName)

	//multiple deklarasi const

	const (
		name     string = "amin"
		lastName        = "jaya"
	)
	fmt.Println(name, lastName)
}
