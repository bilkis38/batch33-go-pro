package main

import (
	"fmt"
	"time"
)

type location struct {
	id       int64
	addresss string
}

type departement struct {
	id             int64
	departmentName string
	location       location
}

type employee struct {
	firstName   string
	lastName    string
	hireDate    time.Time
	salary      float64
	departement departement
}

func toString(e employee) string {
	return fmt.Sprintf("FullName: %s %s, HireDate: %s, Salary: %.2f, Departemen:%s, Address:%s",
		e.firstName, e.lastName, e.hireDate.Format("2025-02-02"), e.salary, e.departement.departmentName, e.departement.location.addresss)
}

func main() {
	//1. literal
	emp1 := employee{
		firstName: "Bilkis",
		lastName:  "Nisa",
		hireDate:  time.Now(),
		salary:    150_000,
		departement: departement{
			id:             10,
			departmentName: "IT",
			location: location{
				id:       1001,
				addresss: "Jakarta",
			},
		},
	}
	fmt.Println(toString(emp1))
}
