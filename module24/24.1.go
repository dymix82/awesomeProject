// Задание 1. Сортировка вставками
// Что нужно сделать
// Напишите функцию, сортирующую массив длины 10 вставками.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 10

func main() {
	ar := make([]int, size)
	for k := 0; k < len(ar); k++ { // Наполняем массив
		rand.Seed(time.Now().UnixNano())
		time.Sleep(2)
		x := rand.Intn(20)
		//		fmt.Println(x)
		ar[k] = x
	}
	InsertionSort(ar)
	fmt.Println("Отсортированный массив:")
	fmt.Println(ar)
}

func InsertionSort(ar []int) {
	for i := 1; i < len(ar); i++ {
		x := ar[i]
		j := i
		for ; j >= 1 && ar[j-1] > x; j-- {
			ar[j] = ar[j-1]
		}
		ar[j] = x
	}
}
