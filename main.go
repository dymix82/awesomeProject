package main

import (
	"fmt"
)

func main() {
	var n1, n2, n3, n4, n5, n6, result int
	for i := 100000; i <= 999999; i++ {
		n1 = i / 100000
		n2 = (i / 10000) % 10
		n3 = (i / 1000) % 10
		n4 = (i / 100) % 10
		n5 = (i / 10) % 10
		n6 = i % 10
		//fmt.Println(n1,n2,n3,n4,n5,n6)
		if (n1 != n6) && (n2 != n5) && (n3 != n4) {
			continue
		}
		result++
	}
	fmt.Println("Количество зеркальных билетов среди всех шестизначных номеров")
	fmt.Println("от 100000 до 999999:")
	fmt.Println(result)
}
