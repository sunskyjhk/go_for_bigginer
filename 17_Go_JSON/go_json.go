package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

func main() {
	fmt.Println("Hello World")

	author := Author{"Jihoon", 35, true}
	book := Book{"Learning Concurrency in Python", author}

	fmt.Printf("%+v\n", book)

	byteArray, err := json.MarshalIndent(book, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
}
