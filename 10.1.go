// Напишите программу, вычисляющую ex посредством разложения в ряд
// Тейлора с заданной пользователем точностью.
// Пользователь вводит значение x и с точностью до какого знака
// после запятой необходимо вычислить значение функции.
package main

import (
	"fmt"
	"math"
)

func main() {
	var exp float64 = 1.0
	var prev float64
	var x, acc, n int
	fmt.Println("Введите значение x и точность до какого знача после запятой необходимо вычислить значение функции")
	fmt.Scan(&x)
	fmt.Scan(&acc)
	epsilon := 1 / math.Pow(10, float64(acc))
	var factarial float64 = 1.0
	n = 1

	for {
		factarial *= float64(n)
		prev = exp
		exp += math.Pow(float64(x), float64(n)) / float64(factarial)
		if math.Abs(exp-prev) < epsilon {
			break
		}
		n++

	}
	fmt.Println(exp)
}
