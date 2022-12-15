package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func squares(c chan int) {

	for {
		num := <-c
		fmt.Println(num * num)
	}
}

func main() {
	c := make(chan int)
	s := make(chan os.Signal, 1)
	go squares(c)
	signal.Notify(s, os.Interrupt, syscall.SIGTERM)

	for i := 1; ; i++ {
		select {
		case <-s:
			fmt.Println("выхожу из программы")
			return
		default:
			c <- i
		}
	}

}
