package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTPKiki NoKTP = "1234567890"
	var marriedStatus Married = false
	
	fmt.Println(noKTPKiki)
	fmt.Println(marriedStatus)


}