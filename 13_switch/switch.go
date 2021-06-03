package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "domingo"
	case 2:
		return "segunda"
	case 3:
		return "terça"
	case 4:
		return "quarta"
	case 5:
		return "quinta"
	case 6:
		return "sexta"
	case 7:
		return "sabádo"
	default:
		return "numero invalido"
	}
}

func diaDaSemana2(numero int) string {

	switch {
	case numero == 1:
		return "domingo"
	case numero == 2:
		return "segunda"
	case numero == 3:
		return "terça"
	case numero == 4:
		return "quarta"
	case numero == 5:
		return "quinta"
	case numero == 6:
		return "sexta"
	case numero == 7:
		return "sábado"
	default:
		return "número inválido"
	}

}

func main() {
	fmt.Println("switch")
	fmt.Println("------")

	dia := diaDaSemana(500)
	fmt.Println(dia)
}
