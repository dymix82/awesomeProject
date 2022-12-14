package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var file1 string
	var file2 string
	var file3 string
	var Result string
	flag.StringVar(&file1, "file1", "null", "str")
	flag.StringVar(&file2, "file2", "null", "substr")
	flag.StringVar(&file3, "result", "null", "substr")
	flag.Parse()
	strFile1, err := os.ReadFile(file1)
	if err != nil {
		fmt.Print(err)
	}
	strFile2, err := os.ReadFile(file2)
	if err != nil {
		fmt.Print(err)
	}
	if file2 != "null" {
		ResArray := []string{string(strFile1), string(strFile2)}
		Result = strings.Join(ResArray, "\n")
	} else {
		Result = string(strFile1)
	}
	if file3 != "null" {
		f, err := os.Create(file3)
		if err != nil {
			fmt.Print(err)
		}
		defer f.Close()
		_, err2 := f.WriteString(Result)
		if err2 != nil {
			fmt.Print(err2)
		}
	} else {
		fmt.Println(Result)
	}

}
