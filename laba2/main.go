package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//region #1 задание
	isEven(1) //определение четное , нечетное число
	isEven(2)

	//region #2 задание
	Exercise2(-1)
	Exercise2(2) //позитив , негатив , нноль
	Exercise2(0)

	//region #3 задание
	for i := 1; i <= 10; i++ { //цикл от 1 до 10
		fmt.Println(i)
	}

	//region #4 задание
	getLen("Приветик") // фун-я которая принимает строку и выводит ее длину
	getLen("Я очень люблю веб-разработку")

	//region #5 задание
	rect1 := Rectangle{5, 2}
	rect1.GetRectanglePloshad() // создать структру и реализовать метод для вычисления площади прямоугольника

	//region #6 задание
	//Принимает любое количество значений
	avgNums(2, 3) 
}

func avgNums(num ...int) {
	sum := 0
	counter := 0                 // ф-я для 6-го задания
	for _, v := range num {
		counter++
		sum += v
	}
	fmt.Println(float64(sum) / float64(counter))
}

type Rectangle struct {
	osnovanie int
	visota    int
}

func (r Rectangle) GetRectanglePloshad() {
	fmt.Println((float64(r.osnovanie) * float64(r.visota)) )
}     // вычисление площади

// Русские символы считаются за 2 при обычном len
func getLen(str string) {
	fmt.Println(utf8.RuneCountInString(str))
} // функция для 4-го задания

func isEven(num int) {
	if num%2 == 0 {
		fmt.Println("Чётное")   // функция для 1-го задания
	} else {
		fmt.Println("Нечётное")
	}
}

func Exercise2(num int) {
	if num > 0 {
		fmt.Println("Positive")
	} else if num < 0 {           // функция для 2-го задания
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}

}
