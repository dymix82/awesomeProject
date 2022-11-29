// Задание 1. Функция, возвращающая результат
// Что нужно сделать
// Напишите функцию, которая на вход получает число и возвращает true,
// если число четное, и false, если нечётно
package main

import "fmt"

func main() {
	var a int
	fmt.Println("----------\nВведите число\n----------")
	fmt.Scan(&a)
	fmt.Println(chetnechet(a))
}
func chetnechet(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}
