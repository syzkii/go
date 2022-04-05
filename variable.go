package main

import "fmt"

func main(){

//simpel use
	var (
		firstName = "Kiki"
		lastName =  "Rifki"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
//tanpa menyebutkan tipe data string, jadi lebih simpel menggunakan ":=" 
	country := "Indonesia"
	fmt.Println(country)
//menggunakan string dirasa kurang simpel
	var nama string

	nama = "kiki"
	fmt.Println(nama)

	var name = "Rifki"
	fmt.Println(name)

}