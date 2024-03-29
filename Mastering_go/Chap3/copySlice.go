package main

import "fmt"

func main() {
	a6 := []int{-10, 1, 2, 3, 4, 5} // len() =>6
	a4 := []int{-1, -2, -3, -4}     // len() => 4
	fmt.Println("a6 :", a6)
	fmt.Println("a4 :", a4)

	copy(a6, a4)
	// 배열의 첫번째 부분 부터 덮어 버림
	fmt.Println("a6 :", a6)
	fmt.Println("a4 :", a4)
	fmt.Println()

	b6 := []int{-10, 1, 2, 3, 4, 5}
	b4 := []int{-1, -2, -3, -4}
	fmt.Println("b6 :", b6)
	fmt.Println("b4 :", b4)

	copy(b6, b4)
	fmt.Println("b6 :", b6)
	fmt.Println("b4 :", b4)
	fmt.Println()

	array4 := [4]int{4, -4, 4, -4}
	s6 := []int{1, 1, -1, -1, 5, -5}
	copy(s6, array4[0:])
	// 슬라이스로 변환
	fmt.Println("array4 :", array4[0:])
	fmt.Println("s6 :", s6)
	fmt.Println()

	array5 := [5]int{5, -5, 5, -5, 5}   //len() => 5
	s7 := []int{7, 7, -7, -7, 7, -7, 7} //len() =>7
	copy(array5[0:], s7)
	fmt.Println("array5 :", array5)
	fmt.Println("s7 :", s7)
}
