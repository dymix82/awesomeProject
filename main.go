package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введие число 1:")
	var a int
	fmt.Scan(&a)
	fmt.Println("Введие число 2:")
	var b int
	fmt.Scan(&b)
	var result int
	result = a % b
	if result == 0 {
		fmt.Println("Число делится без остатка")
	} else {
		fmt.Println("Число делится с остатком")
	}
}
