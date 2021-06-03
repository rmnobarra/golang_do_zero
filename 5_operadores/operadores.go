package main

import "fmt"

func main() {
	//aritmeticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 2
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//fim aritiméticos

	//atribuição
	var variavel1 string = "string"
	variavel2 := "String 2"

	fmt.Println(variavel1, variavel2)

	//fim operadores atribuição

	//operadores relacionais

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//fim operadores relacionais

	// operadores lógicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) // and
	fmt.Println(verdadeiro || falso) // or
	fmt.Println(!verdadeiro)         //not
	fmt.Println(!falso)              //not

	//fim operadores lógicos

	// operadores unários
	numero := 10
	numero++
	numero + 15 //numero = numero + 15
	numero--
	numero -= 20 //numero = numero - 20
	numero *= 3  //numero = numero * 3
	numero /= 3
	numero %= 3
	fmt.Println(numero)

	//fim operadores unarios

	//operador ternario
	var texto string
	if numero > 5 {
		texto = "maior que 5"
	} else {
		texto = "menor que 5"
	}
	fmt.Println(texto)
}
