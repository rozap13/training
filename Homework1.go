package main

import (
	"fmt"
	"strconv"
)

func math() {

	var num1 float64
	var num2 float64
	var operator string
	var res float64
	var operatorName string

	fmt.Println("Введите первое число")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		fmt.Println("Не удалось считать значение!")
	}
	fmt.Println("Введите арифмитический оператор")
	_, err2 := fmt.Scanln(&operator)
	if err2 != nil {
		fmt.Println("Не удалось считать значение!")
	}
	fmt.Println("Введите второе число")
	_, err3 := fmt.Scanln(&num2)
	if err3 != nil {
		fmt.Println("Не удалось считать значение!")
	}

	switch operator {
	case "+":
		res = num1 + num2
		operatorName = "сложения"
	case "-":
		res = num1 - num2
		operatorName = "вычитания"
	case "*":
		res = num1 * num2
		operatorName = "умножения"
	case "/":
		res = num1 / num2
		operatorName = "деления"
	}
	fmt.Printf("Результат %s: %s", operatorName, strconv.FormatFloat(res, 'f', -1, 64))
}
