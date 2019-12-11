package main

import "fmt"

func main() {
	/*
		aMap := map[string]int{}
		aMap["test"] = 1
	*/

	aMap := map[string]int{}
	aMap = nil
	fmt.Println(aMap)
	aMap["test"] = 1
	// nil 맵에는 데이터를 넣을 수 없다.
}
