package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("log.txt")

	if err != nil && errors.Is(err, os.ErrNotExist) {
		fmt.Println("Файла не сущестует")
		return
	}
	defer f.Close()

	fi, err := os.Stat("log.txt")
	if err != nil {

	}
	size := fi.Size()
	buf := make([]byte, size)
	if _, err := io.ReadFull(f, buf); err != nil {
		panic(err)
	}

	fmt.Println(size)
	fmt.Printf("%s\n", buf)

}
