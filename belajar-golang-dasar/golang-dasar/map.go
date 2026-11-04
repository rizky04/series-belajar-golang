package main
import "fmt"

func main(){
	
	person := map[string]string{
		"name": "eko",
		"address": "subang",
	}
	fmt.Println(person)

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:",len(m))

	delete(m, "k2")
	fmt.Println("map", m)

	book := make(map[string]string)
	book["title"] = "buku Golang"
	book["author"] = "eko kurniawan"
	book["wrong"] = "ups"

	

	delete(book, "wrong")
	fmt.Println(book)
}