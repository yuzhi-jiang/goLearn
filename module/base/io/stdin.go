package main

import (
	"bufio"
	"fmt"
	"os"
)

func main1() {

	var (
		firstName, lastName, s string
		i                      int
		f                      float32
		input                  = "56.12 / 5212 / Go"
		format                 = "%f / %d / %s"
	)
	fmt.Print("enter your  full name:")
	fmt.Scanln(&firstName, &lastName)

	fmt.Printf("you full name is %s  %s\n", firstName, lastName)

	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("from the string we read: ", f, i, s)

}

/*
*
缓存读取
*/
func main2() {
	var inputReader *bufio.Reader
	var input string
	var err error
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}

}

/*
*
case使用
*/
func main3() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	fmt.Printf("Your name is %s", input)
	// For Unix: test with delimiter "\n", for Windows: test with "\r\n"
	switch input {
	case "Philip\r\n":
		fmt.Println("Welcome Philip!")
	case "Chris\r\n":
		fmt.Println("Welcome Chris!")
	case "Ivo\r\n":
		fmt.Println("Welcome Ivo!")
	default:
		fmt.Printf("You are not welcome here! Goodbye!")
	}

	// version 2:
	switch input {
	case "Philip\r\n":
		fallthrough
	case "Ivo\r\n":
		fallthrough
	case "Chris\r\n":
		fmt.Printf("Welcome2 %sPhilip", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}

	// version 3:
	switch input {
	case "Philip\r\n", "Ivo\r\n":
		fmt.Printf("Welcome1 %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}
}

func main() {
	main3()
}
