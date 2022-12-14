// Задание 2. Анонимные функции
//Что нужно сделать
//Напишите функцию, которая на вход принимает функцию вида A func (int, int) int,
//а внутри оборачивает и вызывает её при выходе (через defer).
//
//Вызовите эту функцию с тремя разными анонимными функциями A.
//Тела функций могут быть любыми, но главное, чтобы все три выполняли разное действие.

package main

import (
	"fmt"
)

func main() {
	anon(4, 3, func(x, y int) int { return x * y })
	anon(3, 4, func(x, y int) int { return x - y })
	anon(4, 3, func(x, y int) int { return x / y })
}
func anon(x int, y int, A func(x, y int) int) {
	defer fmt.Println(A(x, y))
}
