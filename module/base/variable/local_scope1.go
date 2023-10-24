package main

var b = "G"

func main() {
	n1()
	m1()
	n1()
}

func n1() { print(b) }

func m1() {
	b = "O"
	print(b)
}
