package main

import "fmt"

func main() {
	fizzBuzz(15)
	fmt.Println("")
	fizz, buzz, fizzBuzz := fizzBuzzCount(15)
	fmt.Printf("fizz:%d Buzz:%d Count:%d", fizz, buzz, fizzBuzz)
	fmt.Println("")
	x, y, z := fizzBuzzSum(15)
	fmt.Printf("Sumfizz:%d SumBuzz:%d SumCount:%d", x, y, z)
}

func fizzBuzzSum(n int) (int, int, int) {
	fizz := 0
	buzz := 0
	fizzBuzz := 0
	for i := 1; i <= n; i++ {
		// fmt.Printf("%d ", i)
		if i%3 == 0 && i%5 == 0 {
			fizzBuzz += i
		} else if i%3 == 0 {
			// fizz=fizz+1
			// fizz = fizz+i
			fizz += i
		} else if i%5 == 0 {
			buzz += i
		}
	}
	return fizz, buzz, fizzBuzz
}

func fizzBuzzCount(n int) (int, int, int) {
	fizz := 0
	buzz := 0
	fizzBuzz := 0
	for i := 1; i <= n; i++ {
		// fmt.Printf("%d ", i)
		if i%3 == 0 && i%5 == 0 {
			fizzBuzz++
		} else if i%3 == 0 {
			// fizz=fizz+1
			fizz++
		} else if i%5 == 0 {
			buzz++
		}
	}
	return fizz, buzz, fizzBuzz
}

func fizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		// fmt.Printf("%d ", i)
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%s", "fizzBuzz")
		} else if i%3 == 0 {
			fmt.Printf("%s ", "fizz")
		} else if i%5 == 0 {
			fmt.Printf("%s ", "buzz")
		} else {
			fmt.Printf("%d ", i)
		}
	}

}
