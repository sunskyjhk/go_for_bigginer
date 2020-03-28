package main

import "fmt"

func main() {
	var amazing bool
	amazing = true
	if amazing {
		subscribeToChannel()
	}
}

func subscribeToChannel() {
	fmt.Println("this is amazing!")
}
