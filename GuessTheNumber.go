package main

import (
	"fmt"
	"math/rand"
)

func GuessTheNumber() {
	var number int
	var thatNumber int
	var answer string
	attempts := 0
	thatNumber = rand.Intn(100) + 1
	fmt.Println("Угадайте число от 0 до 100!")
	for {
		attempts++
		fmt.Println("введите число")
		_, err := fmt.Scan(&number)
		if err != nil {
			fmt.Println(err)
		}
		if number < thatNumber {
			fmt.Println("Ваша число меньше загаданного")
		}
		if number > thatNumber {
			fmt.Println("Ваша число больше загаданного")
		}
		if number == thatNumber {
			fmt.Printf("Поздравляю! Вы угадали за %d попыток! \n", attempts)
			fmt.Println("Хотите попробовать снова? Да/Нет")
			_, err = fmt.Scan(&answer)
			if err != nil {
				fmt.Println(err)
				break
			}
			if answer != "Да" {
				break
			} else {
				thatNumber = rand.Intn(100) + 1
			}
		}
		if attempts > 100 {
			break
		}
	}
}
