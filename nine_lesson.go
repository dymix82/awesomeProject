package main

import (
	"fmt"
	"math"
)

func main() {
	MaxUint8 := math.MaxUint8
	MaxUint16 := math.MaxUint16
	MaxUint32 := math.MaxUint32
	CountUint8 := 0
	CountUint16 := 0
	fmt.Println(MaxUint8)
	for i := MaxUint32; i >= MaxUint8; i = i - MaxUint8 {
		CountUint8++
	}
	for i := MaxUint32; i >= MaxUint16; i = i - MaxUint16 {
		CountUint16++
	}

	fmt.Println("Количество переполнений uint8 составляет", CountUint8, "в диапазоне от 0 до uint32")
	fmt.Println("Количество переполнений uint16 составляет", CountUint16, "в диапазоне от 0 до uint32")
}
