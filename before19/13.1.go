package main

import "fmt"

func summ(p, q int) int {
	count := 0
	var mid int
	if p > q {
		mid = p
		p = q
		q = mid
	}
	for i := p + 1; i < q; i++ {
		if i%2 == 0 {
			count += i
		}
	}
	return count
}
func main() {

	var myvar1 = summ(2, 10)

	fmt.Printf("\nСумма четных чисел между двумя числами равна %d", myvar1)
}
