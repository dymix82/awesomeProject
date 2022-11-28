// Задание 1. Подсчёт определителя
//Что нужно сделать
//Напишите функцию, вычисляющую определитель матрицы размером 3 × 3.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 3

func fill() [size][size]int { // Функция наполнения матрицы случайными числами от 0 до 9
	A := [size][size]int{}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			rand.Seed(time.Now().UnixNano())
			time.Sleep(2)
			x := rand.Intn(10)
			A[i][j] = x
		}
	}
	return A
}
func determinat(A [size][size]int) int { // Вычисляем дискриминант
	D := (A[0][0] * A[1][1] * A[2][2]) - (A[0][0] * A[1][2] * A[2][1]) - (A[0][1] * A[1][0] * A[2][2]) + (A[0][1] * A[1][2] * A[2][0]) + (A[0][2] * A[1][0] * A[2][1]) - (A[0][2] * A[1][1] * A[2][0])
	return D
}

func main() {
	matrix := [size][size]int{}
	matrix = fill()
	fmt.Println("Ищем определитель матрицы:")
	for i := 0; i < size; i++ { // Циклом печатаем матрицу 3 на 3 в удобном виде
		fmt.Println(matrix[i])
	}
	Det := determinat(matrix)                // Присваиваем значение определителя матрицы
	fmt.Printf("Определитель равен %d", Det) // Выводим результат
}
