package main

import "fmt"

func main() {
	var array1 [6]int = [6]int{1, 2, 3, 4, 5, 6}
	var array2 = [6]int{1, 2, 3, 4, 5, 6}
	var array3 = [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println("array1:", array1)
	fmt.Println("array2:", array2)
	fmt.Println("array3:", array3)

	zeros := [8]int{}
	ptrs := [8]*int{}
	emptystr := [8]string{}

	fmt.Println("zeros:", zeros)
	fmt.Println("ptrs:", ptrs)
	fmt.Println("emptystr:", emptystr)

	type Data struct {
		Number int
		Text   string
	}

	structs := [8]Data{}

	fmt.Println("structs:", structs)
}
