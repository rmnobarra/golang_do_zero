package main

import "fmt"

var n int

func init() {
	fmt.Println("Executando função init")
	n = 20
}

func main() {
	fmt.Println("Função init")
	fmt.Println(n)
}
