package main

import "fmt"

func main() {
	// converting from in to float
	var myint int
	myint = 32
	myfloat := float64(myint)
	fmt.Println(myfloat)

	// converting from float to int
	var myfloat2 float64
	myfloat2 = 3.14
	myint2 := int(myfloat2)
	fmt.Println(myint2)
}
