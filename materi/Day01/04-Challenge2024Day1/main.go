package main

import "fmt"

//box start
func main() {
	boxStar(5)
	boxHollow(5)
	fmt.Println("")
	triangle01(5)
	fmt.Println("")
	triangel02(5)
	fmt.Println("")
	triangel03(5)
	fmt.Println("")
	triangle04(5)
	fmt.Println("")
	diagonal(5)
	fmt.Println()
	elPattern(5)
	fmt.Println()
	var count = countDigit(12345678)
	fmt.Println("count digit:", count)
	fmt.Println("count digit2:", countDigit(12345678))

	fmt.Printf("is Prime : %t", isPrime(5))

	fmt.Println()
	fmt.Printf("is Prime : %t", isPrime(12))

	for i := 2; i <= 100; i++ {
		if isPrime(i) {
			fmt.Printf("%d", i)
		}
	}
}
func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countDigit(n int) int {
	count := 0
	sisa := n

	for {
		sisa = sisa / 10
		count++
		if sisa == 0 {
			break
		}
	}
	return count
}

func elPattern(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j == 1 {
				fmt.Printf("%d ", i)
			} else if i == n {
				fmt.Printf("%d ", i+j-1)
			} else {
				fmt.Printf("%s ", "*")
			}
		}
		fmt.Println("")
	}
}

func diagonal(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j {
				fmt.Printf("%d ", i)
			} else {
				fmt.Printf("%s ", "*")
			}
		}
		fmt.Println("")
	}
}

func triangle04(n int) {
	for i := 1; i <= n; i++ {
		for j := n; j >= 1; j-- {
			if j <= (n+1)-i {
				fmt.Printf("%s ", "*")
			} else {
				fmt.Printf("%s ", " ")
			}

		}
		fmt.Println("")
	}
}

func triangel03(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j <= n-i {
				fmt.Printf("%s ", " ")
			} else {
				fmt.Printf("%s ", "*")
			}
		}
		fmt.Println("")
	}
}

func triangel02(n int) {
	for i := 1; i <= n; i++ {
		for j := (n + 1) - i; j >= 1; j-- {
			fmt.Printf("%s ", "*")
		}
		fmt.Println("")
	}
}

func triangle01(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%s ", "*")
		}
		fmt.Println("")
	}
}

func boxHollow(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 || i == n || j == 1 || j == n {
				fmt.Printf("%s ", "*")
			} else {
				fmt.Printf("%s ", " ")
			}

		}
		fmt.Println("")
	}

}

//box start
func boxStar(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%s ", "*")
		}
		fmt.Println("")
	}
}
