package main

import (
	"fmt"
)

type Employee struct {
	ID     int
	Salary int
	Name   string
}

var employeeList []Employee
var nameList [3]string

func EmployeeById(id int) *Employee {
	fmt.Printf("id: %v\n", id)
	for i, _ := range employeeList {
		if employeeList[i].ID == id {
			return &employeeList[i]
		}
	}
	return &Employee{}
}
func main4_4() {
	nameList = [...]string{"Elsa", "Ana", "Aaron"}
	employeeList = []Employee{}
	for i := 0; i < 3; i++ {
		employee := Employee{i, i * 100, nameList[i]}
		employeeList = append(employeeList, employee)
	}

	fmt.Println(EmployeeById(1).Name)
	fmt.Println(EmployeeById(2).Salary)
	EmployeeById(2).Salary = -100
	fmt.Println(EmployeeById(2).Salary)
}
