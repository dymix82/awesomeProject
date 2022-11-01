package main

import "fmt"

func main() {
	var x, count int
	count = 2
	x = -10
	for {
		if count%2 == 0 {
			x++
			fmt.Println(x)
		} else {
			x--
			fmt.Println(x)
		}
		if x == 10 || x == -10 {
			count++
		}
		if x == 0 {
			fmt.Println("Вы перешли КПП")
		}
	}

}
