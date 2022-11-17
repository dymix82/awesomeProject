package main

import (
	"fmt"
	"math"
)

func main() {
	MaxUint8 := math.MaxInt8
	MaxUint16 := math.MaxInt16
	MaxUint32 := math.MaxInt32
	MinInt8 := math.MinInt8
	MinInt16 := math.MinInt16
	var a, b int16
	var c int
	fmt.Println("Введите два числа для определения типа данных их произведения")
	fmt.Scan(&a)
	fmt.Scan(&b)
	c = int(a) * int(b)
	fmt.Println(c)
	if c > 0 && c <= MaxUint8 {
		fmt.Println("результат uint8")
	} else if c > 0 && c <= MaxUint16 {
		fmt.Println("результат uint16")
	} else if c > 0 && c <= MaxUint32 {
		fmt.Println("результат uint32")
	} else if c < MinInt16 {
		fmt.Println("результат int32")
	} else if c > MinInt16 && c < MinInt8 {
		fmt.Println("результат int16")
	} else {
		fmt.Println("результат int8")
	}
	fmt.Println(c)

}
