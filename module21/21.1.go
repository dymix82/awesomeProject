// Задание 1. Расчёт по формуле
//Что нужно сделать
//Напишите функцию, производящую следующие вычисления.
//
//S = 2 × x + y ^ 2 − 3/z, где x — int16, y — uint8, a z — float32.
//
//Тип S должен быть во float32.

package main

import "fmt"

var (
	S float32
	x int16
	y uint8
	z float32
)

func main() {
	x = 4                                            // задаем значение х
	y = 4                                            // задаем значение y
	z = 1.56                                         // задаем значение z
	F := func(x int16, y uint8, z float32) float32 { // Описываем анонимную функцию для нахождения  S = 2 × x + y ^ 2 − 3/z
		S = float32(2*x) + float32(y*y) - (3 / z)
		return S
	}
	S = F(x, y, z) // Вызываем функцию со значениями x,y,z
	fmt.Println(S) // Выводим результат
}
