package main

import "fmt"

type Power4 int

func main() {
	const (
		P4_0 Power4 = 4 << iota
		P4_1
		P4_2
		P4_3
		P4_4
	)
	fmt.Println(P4_0, P4_1, P4_2, P4_3, P4_4)
}
