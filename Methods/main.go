package main

import "fmt"

type Employee struct {
	Name       string
	Id         int
	Experience int
}

func (e Employee) getId() {
	fmt.Println(e.Id)
}

func (e Employee) getAll() {
	fmt.Println(e.Name, e.Id, e.Experience)
}
func main() {

	emp1 := Employee{
		Name:       "E1",
		Id:         1,
		Experience: 2,
	}
	fmt.Println(emp1)
	emp1.getId()
	emp1.getAll()
}
