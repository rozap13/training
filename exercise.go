package main

import (
	"fmt"
)

func main() {
	// Ваш код помещайте после этой строки и до символа }

	var number int
	fmt.Println("Введите целое число от 1 до 100")
	fmt.Scan(&number)
	if number > 100 || number < 1 {
		fmt.Print("Указано некорректное число!")
	} else {
		for i := 1; i <= number; i++ {
			if isPrime(i) {
				if i > 2 {
					fmt.Print(", ")
				}
				fmt.Print(i)
			}
		}
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i < n; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
