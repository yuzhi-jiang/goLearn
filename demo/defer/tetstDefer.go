package main

import "fmt"

type student struct {
	age int
}

func main() {
	stud1 := student{1}
	stud1.age = 2
	defer printVar(&stud1)
	stud1.age = 3
	fmt.Println(stud1.age)

}

func printVar(a *student) {
	//print address
	println("before:", a.age)
	a.age = 4
	println("after:", a.age)
}
