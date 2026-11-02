package main 
import "fmt"

func main(){
names := [...]string{"eko","kurniawan","joko","anwar","budi","nugraha","jaya"}

slice1 := names[4:6]
fmt.Println(len(slice1))

slice2 := names[:3]
fmt.Println(cap(slice2))

slice3 := names[3:]
fmt.Println(slice3)

slice4 := names[0]
fmt.Println(slice4)

days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
daySlice1 := days[5:]//diambil slice dari data array ke 5
fmt.Println(daySlice1)
daySlice1[0] = "Sabtu Baru"//diganti slicenya
daySlice1[1] = "Minggu Baru"//diganti slicenya
fmt.Println(days)

daySlice2 := append(daySlice1, "Libur Baru")
daySlice2[0] = "Okes"
fmt.Println(daySlice2)
fmt.Println(days)


var newSlice []string = make([]string, 2,5)
newSlice[0] = "eko"
newSlice[1] = "eko"
//newSlice[2] = "eko" harusnya menggunakan append

fmt.Println(newSlice)
fmt.Println(len(newSlice))
fmt.Println(cap(newSlice))

newSlice2 := append(newSlice, "eko")
fmt.Println(newSlice2)
fmt.Println(len(newSlice2))
fmt.Println(cap(newSlice2))

newSlice2[0] = "budi"
fmt.Println(newSlice2)
fmt.Println(newSlice)


fromSlice := days[:]
toSlice := make([]string, len(fromSlice), cap(fromSlice))
copy(toSlice, fromSlice)
fmt.Println(fromSlice)
fmt.Println(toSlice)
}