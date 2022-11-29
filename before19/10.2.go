// В связи с особенностями хранения дробных чисел они не очень подходят для хранения
// денежных значений (может потеряться значимая часть при переполнении, да и дробная
//
//	часть больше двух знаков не нужна).
//
// Но мы попробуем решить задачу начисления процентов, используя именно их.
package main

import (
	"fmt"
	"math"
)

func main() {
	var AmountOfDeposit, TimeOfDeposit int
	var vklad, vkladokr, Proc, sdacha float64
	fmt.Println("Введите сумму вклада")
	fmt.Scan(&AmountOfDeposit)
	fmt.Println("Введите срок хранения (года) ")
	fmt.Scan(&TimeOfDeposit)
	fmt.Println("Процент")
	fmt.Scan(&Proc)
	months := 12 * TimeOfDeposit
	vklad = float64(AmountOfDeposit)
	monthproc := float64(Proc / 100)
	for i := 1; i <= months; i++ {
		vklad = vklad + vklad*monthproc
		vkladokr = math.Floor(vklad*100) / 100
		sdacha += (vklad - vkladokr)
	}
	fmt.Println(vkladokr, sdacha)
}
