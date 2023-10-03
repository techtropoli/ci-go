package main

import "fmt"

func main() {
	fmt.Println(soma(10, 10))
	fmt.Println(subtracao(10, 10))
}

func soma(a int, b int) int {
	return a + b
}

func subtracao(a int, b int) int {
	return a - b
}
