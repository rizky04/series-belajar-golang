package main 
import "fmt"

func main(){
	// names := make(map[string]string)
	// names["name1"] = "joko"
	// names["name2"] = "anwar"
	// names["name3"] = "dolor"
	// name := "eko"
	var name = [3]string{
		"eko",
		"anwar",
		"widodo",
	}
	fmt.Println(name)
	if name == "eko"{
		fmt.Println("Hello")
	}else if name == "joko"{
		fmt.Println("hi boleh kenalan")
	}else{
		fmt.Println("gak mau kenala")
	}
}