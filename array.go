package main

import "fmt"

func main() {

	//var names [4] string
	//names[0] = "Rifki"
	//names[1] = "Ardian"
	//names[2] = "Akbar"
	//names[3] = "aaa"

	//fmt.Println(names[0])
	//fmt.Println(names[1])
	//fmt.Println(names[2])
	//fmt.Println(names[3])

	var values = [3] int{
		11,
		112,
		223,
	}
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values)) // len = panjang datanya bukan jumlah datanya
	fmt.Println(values)
}