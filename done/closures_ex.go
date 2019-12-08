package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}

/*
아직 이해가 부족함
TODO:
	https://manducku.tistory.com/56
	https://steemit.com/golang/@jungmu/a-tour-of-go-43-function-closures
	https://www.calhoun.io/5-useful-ways-to-use-closures-in-go/
*/
