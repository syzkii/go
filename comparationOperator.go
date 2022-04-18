package main

import "fmt"

func main() {

	var name1 = "Kiki"
	var name2 = "Kiki"

	var result bool = name1 == name2 // true : karena kedua nama string sama
	fmt.Println(result)

	var value1 = 200
	var value2 = 100

	fmt.Println(value1 > value2) // true : value1 lebih dari value2
	fmt.Println(value1 < value2) // false : value1 kurang dari value2
	fmt.Println(value1 == value2) // false : value1 sama dengan value2
	fmt.Println(value1 != value2) // true : value1 tidak sama dengan value2

}