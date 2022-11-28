// Задание 1. Слияние отсортированных массивов
// Что нужно сделать:  Напишите функцию, которая производит слияние
// двух отсортированных массивов длиной четыре и пять в один массив длиной девять.
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func rands() (x int) { // Функция наполнения массива
	rand.Seed(time.Now().UnixNano())
	time.Sleep(2)
	x = rand.Intn(20)
	return x
}

func main() {
	a := make([]int, 5, 5)
	b := make([]int, 4, 4)

	for k := 0; k < len(a); k++ {
		a[k] = rands()
	}
	for k := 0; k < len(b); k++ {
		b[k] = rands()
	}
	fmt.Printf("Исходные массивы: %v и %v \n", a, b)
	sort.Ints(a)
	sort.Ints(b)
	c := append(a, b...)
	fmt.Printf("Исходные отсортированные массивы: %v и %v\n", a, b)
	sort.Ints(c)
	fmt.Printf("Итоговый массив: %v\n", c)

}
