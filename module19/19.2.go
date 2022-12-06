//Задание 2. Сортировка пузырьком
//Что нужно сделать
//Отсортируйте массив длиной шесть пузырьком.
//
//Советы и рекомендации
//Принцип сортировки пузырьком можно посмотреть на «Википедии», там есть наглядная демонстрация, или на YouTube.
//В качестве среды разработки может помочь GoLand или VS Code.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 6

func fill(a []int) (b []int) { // Функция наполнения массива

	return
}
func BubbleSort(a []int) []int {
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	//	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	con := 0
	for i := 0; i < (size)/2; i++ {
		con = a[size-1-i]
		a[size-1-i] = a[i]
		a[i] = con
	}
	return a
}

func main() {
	array1 := make([]int, size)
	for k := 0; k < len(array1); k++ {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(2)
		x := rand.Intn(20)
		//		fmt.Println(x)
		array1[k] = x
	}
	fmt.Println("Вводный массив: ", array1)
	array1_sorted := BubbleSort(array1)
	fmt.Println("Отсортированный массив:", array1_sorted)
}
