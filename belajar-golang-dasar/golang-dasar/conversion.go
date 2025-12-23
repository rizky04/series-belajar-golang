package main

import "fmt"

func main() {
	// Tipe Data | Nilai Minimum | Nilai Maksimum
	// int8 | -128 | 127
	// int16 | -32768 |32767
	// int32 | -2147483648 | 2147483647
	// int64 | -9223372036854775808 | 9223372036854775807

	//jika nilai conversion lebih besar dari nilai diatas 1 angka saja maka minus

	//conversion
	var nilai32 int32 = 32769
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	//conversion tipe data
	var name = "amin"
	var e uint8 = name[1]
	var eString = string(e)
	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

}
