package main

import (
	"fmt"
	"psutiGoLabs/laba3/mathutils"
	"psutiGoLabs/laba3/stringutils"
	"unicode/utf8"
)

func main() {
	//region #1-2 Exercise
	fmt.Println(mathutils.Factorial(5))

	//region #3 Exercise
	stringutils.ReverseString("Katalinka")

	//region #4 Exercise
	arr := make([]int, 5) //заполнение массива от 1 до 5
	for i := 0; i < 5; i++ {
		arr[i] = i + 1
	}
	fmt.Println(arr)

	//region #5 Exercise 
	massiv := [...]int{5, 4, 6, 7, 3} // создаем срез
	//Теперь ссылается на massiv
	newSlice := massiv[:]
	//Ссылка одна и та же, изменения и коснуться massiv
	newSlice[2] = 999 //добавили 
	fmt.Println(massiv, newSlice)

	//После этого слайс превысит len и cap, переаллоцируется в другом месте памяти, больше не будет ссылаться на "massiv"
	newSlice = append(newSlice, 228)
	//Изменения только на переаллоцированном слайсе, не перенеслись как должны были на массив
	newSlice[1] = 1337
	fmt.Println(massiv, newSlice)

	//Удалили последний элемент
	newSlice = newSlice[:len(newSlice)-1]
	fmt.Println(newSlice)

	//region #6 Exercise
	Exercise6("папа", "мама", "акробат", "малина")
}

func Exercise6(strs ...string) { // выполнила срез из строк , и ф-я выберает самую длинную
	strArr := make([]string, len(strs))
	copy(strArr, strs)

	max := ""
	for _, v := range strArr {
		if utf8.RuneCountInString(v) > utf8.RuneCountInString(max) {
			fmt.Println(utf8.RuneCountInString(v))
			max = v
		}
	}

	fmt.Println(max)
}
