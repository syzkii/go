package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 80

	//var lulusUjian = ujian >= 80
	//var lulusAbsensi = absensi >= 80

	//fmt.Println(lulusUjian)
	//fmt.Println(lulusAbsensi)

	//var lulus = lulusUjian && lulusAbsensi
	
	//fmt.Println(lulus)

	fmt.Println(ujian >= 80 && absensi >= 80)
}

// && : dan  true and true = true. true and false = false. false and false = false
// || : atau true or true = true. true or false = true. false or false = false
// ! : Kebalikan jika tidak true = false. jika tidak false = true