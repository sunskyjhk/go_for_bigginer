package main

import (
	"errors"
	"fmt"
)

type Employee interface {
	Language() string
	Age() int
	Random() (string, error)
}

type Engineer struct {
	Name string
}

func (e Engineer) Language() string {
	return e.Name + " programs in Go"
}

func (e Engineer) Age() int {
	return 25
}

func (e Engineer) Random() (string, error) {
	if e.Age() < 20 {
		return "", errors.New("under criterior")
	}
	return "Pass", nil
}

func main() {
	var programmers []Employee
	elliot := Engineer{"Elliot"}
	programmers = append(programmers, elliot)
	fmt.Println(elliot.Language())
	fmt.Println(programmers)
}
