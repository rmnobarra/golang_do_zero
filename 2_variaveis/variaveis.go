package main

import "fmt"

func main() {

	//explicitamente
	var variavel1 string = "variavel 1"

	//implicitamente (inferencia de tipo)
	variavel2 := "variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "variavel 3"
		variavel4 string = "variavel 4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel 5", "variavel 6"

	fmt.Println(variavel5, variavel6)
}
