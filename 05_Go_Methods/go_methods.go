package main

import "fmt"

// Methods

type Employee struct {
	Name string
}

func (e *Employee) UpdateName(newName string) {
	e.Name = newName
}

func (e *Employee) PrintName() {
	fmt.Println(e.Name)
}

/* Functions vs Methods

func UpdateGuitarist(guitarist *Guitarist, params ParamsStruct)  {
	fmt.Println("This is a simple function")
}

// Calling this function
UpdateGuitarist(guitarist, param)

func (g *Guitarist) Update(params ParamsStruct) {
    fmt.Println("This is a simple method")
}

// this is far nicer in my opinion
myGuitarist.Update(params)

*/

func main() {
	// Methods
	var employee Employee
	employee.Name = "Elliot"
	employee.UpdateName("Forbsey")
	employee.PrintName()
}
