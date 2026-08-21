package main
import "fmt"

//operasi &&

//nilai1 | operator | nilai 2 | hasil
//true | && | true | true
//true | && | false | false
//false | && | true | false
//false | && | false | false

//true | || | true | true
//true | || | false | true
//false | || | true | true
//false | || | false | false

//(!) | || | true | false
//(!) | || | false | true
func main(){
 var nilaiAkhir = 90
 var absensi = 81

 var lulusNilaiAkhir bool = nilaiAkhir > 80
 var lulusAbsensi bool = absensi > 80

 var lulus bool =   lulusAbsensi && lulusNilaiAkhir

 fmt.Println(lulus)
}