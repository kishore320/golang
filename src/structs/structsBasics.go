package main

import (
	"fmt"
)

func main() {
	// One way to create struct. Struct literal
	emp1 := Employee{"emp1", "1", 100000}
	// Another way
	emp2 := Employee{
		name:   "emp2",
		id:     "2",
		salary: 110000, // Note ',' is mandatory even for last field
	}
	// Yet another way
	var emp3 Employee
	emp3.name = "emp3"
	emp3.id = "3"
	emp3.salary = 120000

	fmt.Println("Employee one details (name id salary) : ", emp1)
	fmt.Println("Employee two name : ", emp2.name)

	// Even better, use the memory address
	employee := &emp3
	fmt.Println("Employee three salary : ", employee.salary)

	// Slices
	var company []Employee
	company = append(company, emp1, emp2, emp3)

	fmt.Println("Employees in the company are:")
	for index, value := range company {
		fmt.Println(index+1, ":", value)
	}
}

type Employee struct {
	name   string
	id     string
	salary int
}
