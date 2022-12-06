package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var str string
	var substr string
	flag.StringVar(&str, "str", "null", "str")
	flag.StringVar(&substr, "substr", "null", "substr")
	flag.Parse()
	fmt.Println(strings.Contains(str, substr))
}
