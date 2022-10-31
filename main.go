package main

import (
	"fmt"
)

func main() {
    fmt.Println("Введите число 1:")
    var a int
    fmt.Scan(&a)
    fmt.Println("Введие число 2:")
    var b int
    fmt.Scan(&b)
    fmt.Println("Теперь сложите  и введите результат")
    var c int
    fmt.Scan(&c)
    result := a + b 
    if c == result  {
      fmt.Println ("Правильно!" )
      } else {
      fmt.Println ("Неверно!")
      }
}