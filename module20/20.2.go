// Задание 2. Умножение матриц
// Что нужно сделать
// Напишите функцию, умножающую две матрицы размерами 3 × 5 и 5 × 4.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	rowsA = 3 // Кол-во строк матрицы А
	colsA = 5 // Кол-во столбцов матрицы А
	rowsB = 5 // Кол-во строк матрицы В
	colsB = 4 // Кол-во столбцов матрицы В
)

var (
	matrixA [rowsA][colsA]int
	matrixB [rowsB][colsB]int
)

func fillA() [rowsA][colsA]int { // Функция наполнения матрицы A случайными числами от 0 до 9
	A := [rowsA][colsA]int{}
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsA; j++ {
			rand.Seed(time.Now().UnixNano())
			time.Sleep(2)
			x := rand.Intn(10)
			A[i][j] = x
		}
	}
	return A
}
func fillB() [rowsB][colsB]int { // Функция наполнения матрицы B случайными числами от 0 до 9
	A := [rowsB][colsB]int{}
	for i := 0; i < rowsB; i++ {
		for j := 0; j < colsB; j++ {
			rand.Seed(time.Now().UnixNano())
			time.Sleep(2)
			x := rand.Intn(10)
			A[i][j] = x
		}
	}
	return A
}
func MatrixMultiply() [rowsA][colsB]int { // Функция умнажения матрицы А на B
	matrixC := [rowsA][colsB]int{}
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				matrixC[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}
	return matrixC
}
func main() {
	matrixA = fillA()
	matrixB = fillB()
	fmt.Println("Матрица А:")
	for i := 0; i < rowsA; i++ { // Циклом печатаем матрицу A в удобном виде
		fmt.Println(matrixA[i])
	}
	fmt.Println("Матрица B:")
	for i := 0; i < rowsB; i++ { // Циклом печатаем матрицу B в удобном виде
		fmt.Println(matrixB[i])
	}
	fmt.Println("Результат умножения Матрицы А и B:")
	matrixC := MatrixMultiply()  // Применяем функцию умножения матриц
	for i := 0; i < rowsA; i++ { // Циклом печатаем получившуюся матрицу в удобном виде
		fmt.Println(matrixC[i])
	}
}
