package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	printArray(arr1)

	strintArr := []string{"a", "b", "c"}
	printCompareArr(strintArr)
	printStringOrIntArr(strintArr)
}

// 所有泛型
func printArray[T any](arr []T) {
	for _, i := range arr {
		fmt.Println(i)
	}
}

//部分泛型-可以比较的类型（实现了compare）

func printCompareArr[T comparable](arr []T) {
	for _, i := range arr {
		fmt.Println(i)
	}
}

// 部分泛型-string int
func printStringOrIntArr[T string | int](arr []T) {
	for _, i := range arr {
		fmt.Println(i)
	}
}
