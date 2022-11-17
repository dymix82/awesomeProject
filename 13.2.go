package main

import "fmt"

func main() {
	fmt.Println("Введите два числа")
	var n1 int
	var n2 int
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Println("Было:\n------------")
	fmt.Println(n1, n2)
	switchPlaces(&n1, &n2)
	fmt.Println("\nСтало:\n------------")
	fmt.Println(n1, n2)
}

func switchPlaces(n1, n2 *int) {
	*n1, *n2 = *n2, *n1
}
