package main

import (
	"fmt"
	"time"
)

// function untuk menjalankan program
func main() {
	//1, deklarasi variabel dengan data type
	var fullName string = "Jhone Doe"
	fmt.Println(fullName)

	//2. deklarasi multiple variable with datatype using var
	var (
		firstName string    = "jhone"
		lastName  string    = "snow"
		salary    float64   = 5_000_000
		hireDate  time.Time = time.Now()
	)
	fmt.Println(firstName, lastName, salary, hireDate)

	//3. declare multiple variable without datatype
	var (
		departID       = 10
		departmentName = "Finance"
	)
	fmt.Println(departID, departmentName)

	//4. multiple variable in one line with zero velue
	var money1, money2, money3 float64
	fmt.Println(money1, money2, money3)

	//5. short variable with straight value
	var money4 = 56_0000.98
	fmt.Println(money4)

	//6. short variable using :=
	money5 := 45_000.98
	fmt.Println(money5)

	//7. declaring multiple variable with short variable, use :=
	isDiscount, customerName, orderDate, price := false, "sugimoto", time.Now(), 456.89
	fmt.Println(isDiscount, customerName, orderDate, price)

	//8. manipulate short variable
	isDiscount, customerName, orderDate, price = true, "sugimoto", time.Now(), 1000
	fmt.Println(isDiscount, customerName, orderDate, price)
}
