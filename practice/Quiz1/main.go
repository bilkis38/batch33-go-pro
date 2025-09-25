package main

import "fmt"

func main() {
	//Divisor
	var n int
	fmt.Print("Masukkan angka:")
	fmt.Scan(&n)

	fmt.Printf("Bilangan pembagi dari %d adalah : ", n)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Println("", i)

		}
	}
	//ExtractDigit
	var number int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&number)

	fmt.Print("Hasil dibalik: ")
	for number > 0 {
		digit := number % 10 // ambil digit terakhir
		fmt.Printf("%d ", digit)
		number = number / 10 // buang digit terakhir
	}
	fmt.Println()

	//segitiga bintang
	// 3a. var triangel int
	for i := 1; i <= 5; i++ {
		for j := 5; j >= 1; j-- {
			if j <= (5+1)-i {
				fmt.Printf("%s ", "*")
			} else {
				fmt.Printf("%s ", " ")
			}

		}
		fmt.Println("")
	}
}
