package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("Go Linked Lists Tutorial")

	myList := list.New()
	myList.PushBack(1)
	myList.PushFront(2)
	// we now have a linked list with '1' at the back of the list
	// and '2' at the front of the list.

	for element := myList.Front(); element != nil; element = element.Next() {
		fmt.Println(element)
		fmt.Println(element.Value)
		fmt.Println("newList")
	}

	for element := myList.Front(); element != nil; element = element.Next() {
		// do something with element.Value
		if element.Value != 1 {
			myList.Remove(element)
		}
	}

	for element := myList.Front(); element != nil; element = element.Next() {
		// do something with element.Value
		fmt.Println(element.Value)
		fmt.Println("processed list")
	}
}
