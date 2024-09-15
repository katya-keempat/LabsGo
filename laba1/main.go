package main

import (
	"fmt"
	"time"
)

func main() {
	//region #1 задания

	// дата + время
	fmt.Println("Дата и время:")
	fmt.Println(time.Now().Local().Format(time.RFC1123))

	//region #2-3 задания
	someInt, someFloat64, someString, someBool := 13, 12.28, "Katushka", false
	fmt.Println(someInt, someFloat64, someString, someBool) // создание переменных

	//region #4-5 задания
	val1 := doSomeOperation(3, 2, "+") // операции
	fmt.Println(val1)

	val2 := doSomeOperation(3, 2, "-")
	fmt.Println(val2)

	val3 := doSomeOperation(3, 2, "*")
	fmt.Println(val3)

	val4 := doSomeOperation(3, 2, "/")
	
	fmt.Println(val4)   //4задание

	//Проверка на ноль
	val5 := doSomeOperation(3, 0, "/")

	fmt.Println(val5)

	// #5
	sum1, vc1 := Katushka(3.0, 1.8)
	fmt.Println(sum1, vc1)

	//region #6 задание

	var num1 int
	fmt.Scanln(&num1)

	var num2 int
	fmt.Scanln(&num2)

	var num3 int
	fmt.Scanln(&num3)

	fmt.Println(float64(num1+num2+num3) / 3)
}

func doSomeOperation(num1, num2 int, opertaion string) (float64) {
	switch opertaion {
	case "-":
		return float64(num1 - num2)
	case "+":
		return float64(num1 + num2)
	case "*":
		return float64(num1 * num2)
	case "/":
		if num2 != 0 {
			return float64(num1) / float64(num2)
		}
		return -1
	}

	return -1
}

func Katushka(num1, num2 float64) (float64, float64) {
	return num1 + num2, num1 - num2
}