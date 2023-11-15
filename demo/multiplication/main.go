package main

import "fmt"

func main() {

	//Loop through the numbers 1-9
	for i := 1; i <= 9; i++ {
		//Loop through the numbers 1-i
		for j := 1; j <= i; j++ {
			//Print the product of i and j
			fmt.Printf("%d x %d = %d\t", i, j, i*j)
		}
		//Print a new line
		println()
	}

}
