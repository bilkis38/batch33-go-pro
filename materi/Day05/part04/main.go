package main

import (
	"fmt"
	"time"

	"github.com/bootcam/go-pro/Day05/part04/pkg/employee"
)

func main() {

	//1. constructor return pointer employee (recommended)
	//1. pointer employee
	emp1 := employee.NewEmployee("Kang", "Dian", time.Date(2025, time.September, 21, 0, 0, 0, 0, time.UTC), 500_000)
	emp2 := employee.NewEmployee("Rini", "Maharani", time.Date(2025, time.September, 21, 0, 0, 0, 0, time.UTC), 500_000)

	fmt.Println(emp1)
	fmt.Println(emp2)

	//1.update emp1 by emp 3
	emp3 := emp1

	if err := emp3.SetSalary(800_000); err != nil {
		fmt.Println("failed to update salary : ", err)
	}
	emp3.SetFirstName("Asep")
	emp3.SetFirstName("Obor")

	fmt.Println("emp1 after : ", emp1)

	//2. constructor return value employee
	emp4 := employee.NewEmployeeValue("Widi", "Wini", time.Date(2025, time.September, 21, 0, 0, 0, 0, time.UTC), 500_000)
	fmt.Println("emp4 before : ", emp4)
	emp5 := emp4
	if err := emp5.SetSalary(650_000); err != nil {
		fmt.Println("failed to update salary : ", err)
	}
	emp5.SetFirstName("Windi")
	emp5.SetFirstName("Wina")

	fmt.Println("emp4 after : ", emp4)

}
