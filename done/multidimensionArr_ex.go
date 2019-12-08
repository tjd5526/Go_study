package main

import "fmt"

func main() {
	var multiDimArray = [2][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}

	fmt.Println(multiDimArray)
	fmt.Println(multiDimArray[0])
	fmt.Println(multiDimArray[0][1])

	var array = [6]int{1, 2, 3, 4, 5, 6}

	fmt.Println(array[0])
	fmt.Println(array[1])

	array[0] = 6
	fmt.Println(array)

	array[1] = 5
	fmt.Println(array)
}
