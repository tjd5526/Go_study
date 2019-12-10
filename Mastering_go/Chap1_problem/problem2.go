package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		os.Exit(1)
	}
	sum := 0.0
	arguments := os.Args
	for i := 0; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)
		sum += n
	}
	fmt.Println(sum / float64(len(arguments)))
}
