package main

import "fmt"

// floating point numbers definition
var f1 float32
var f2 float64

// define float

var maxFloat32 float32

func main() {
	maxFloat32 = 1688777216
	fmt.Println(maxFloat32 == maxFloat32+10) // you would typically expect this to return false
	// it returns true
	fmt.Println(maxFloat32 + 10)      // 1.6887772e+09
	fmt.Println(maxFloat32 + 2000000) // 1.6907772e+09
}
