package main

import "fmt"

func GetMaximumOfEntered() {
	var number int
	maxnum := 0
	for {
		fmt.Println("Введите число")
		_, err := fmt.Scan(&number)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if number == 0 {
			fmt.Printf("Maximum number is %d", maxnum)
			break
		}
		if number > maxnum {
			maxnum = number
		}
	}
}
