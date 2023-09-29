package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sep := " "
	for i, arg := range os.Args[1:] {
		fmt.Println(strconv.Itoa(i) + sep + arg)
	}
}
