package main

import "fmt"

func main() {
	type NoKTP string

	var ktpEko NoKTP = "3332233"
	fmt.Println(ktpEko)
	fmt.Println(NoKTP("4594949"))
}
