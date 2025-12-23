package main

import "fmt"

func main() {
	// var name string = "amin"
	// fmt.Println(name)
	// name = "ganteng"
	// fmt.Println(name)

	//digolang tidak wajib menggunakan var asalkan kita langsung menginialisasi variabelnya

	// age := 1

	// umur := age + 3
	// fmt.Println(umur)
	// var name = "ganteng"
	// fmt.Println(name)

	//digolang setiap membuat variable harus digunakan jika tidak maka error ketika dijalankan

	//deklarasi multiple var

	var (
		name = "amin"
		age  = 31
	)

	fmt.Println(name, age)
}
