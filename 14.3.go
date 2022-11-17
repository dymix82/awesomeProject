// Задание 3. Именованные возвращаемые значения
// Что нужно сделать
// Напишите программу, которая на вход получает число, затем с помощью двух функций преобразует его.
//
//	Первая умножает, а вторая прибавляет число, используя именованные возвращаемые значения.
package main

import "fmt"

var input int

const a = 5

func main() {
	fmt.Println("----------\nВведите число\n----------")
	fmt.Scan(&input)
	fmt.Println(f1(input))
	fmt.Println(f2(input))
}
func f1(int) (c int) {
	c = input * a
	return
}
func f2(int) (c int) {
	c = input + a
	return
}
