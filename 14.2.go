// Напишите программу, которая с помощью функции генерирует три случайные точки
// в двумерном пространстве (две координаты), а затем с помощью другой функции
// преобразует эти координаты по формулам: x = 2 × x + 10, y = −3 × y − 5.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generatePoint() (x, y int) {
	rand.Seed(time.Now().UnixNano())
	x = rand.Intn(30)
	y = rand.Intn(30)
	return
}
func transform(x, y int) (n, m int) {
	n = 2*x + 10
	m = -3*y - 5
	return
}
func main() {
	x1, y1 := generatePoint()
	fmt.Printf("Координаты точки:\n %d \n %d \n ", x1, y1)
	fmt.Println(transform(x1, y1))
	time.Sleep(1)
	x2, y2 := generatePoint()
	fmt.Printf("Координаты точки:\n %d \n %d \n ", x2, y2)
	fmt.Println(transform(x2, y2))
	time.Sleep(1)
	x3, y3 := generatePoint()
	fmt.Printf("Координаты точки:\n %d \n %d \n ", x3, y3)
	fmt.Println(transform(x3, y3))

}
